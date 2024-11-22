package invoice

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	fileType             = "INVOICES"   // Used by text formats
	magicNumber          = 0x125D       // Used by binary formats
	fileVersion          = 101          // Used by all formats
	dateFormat           = "2006-01-02" // This date must always be used
	nanosecondsToSeconds = 1e9
)

type Invoice struct {
	Id           int
	CustomerId   int
	DepartmentId string
	Raised       time.Time
	Due          time.Time
	Paid         bool
	Note         string
	Items        []*Item
}

type Item struct {
	Id       string
	Price    float64
	Quantity int
	TaxBand  int
	Note     string
}

type InvoicesMarshaler interface {
	MarshalInvoices(writer io.Writer, invoices []*Invoice) error
}

type InvoicesUnmarshaler interface {
	UnmarshalInvoices(reader io.Reader) ([]*Invoice, error)
}

func ProcessInvoices(inFilename, outFilename string, report bool) {
	if inFilename == outFilename {
		log.Fatalln("won't overwrite a file with itself")
	}

	start := time.Now()
	invoices, err := readInvoiceFile(inFilename)
	if err == nil && report {
		duration := time.Now().Sub(start)
		fmt.Printf("Read  %s in %.3f seconds\n", inFilename,
			float64(duration)/nanosecondsToSeconds)
	}
	if err != nil {
		log.Fatalln("Failed to read:", err)
	}
	start = time.Now()
	err = writeInvoiceFile(outFilename, invoices)
	if err == nil && report {
		duration := time.Now().Sub(start)
		fmt.Printf("Wrote %s in %.3f seconds\n", outFilename,
			float64(duration)/nanosecondsToSeconds)
	}
	if err != nil {
		log.Fatalln("Failed to write:", err)
	}
}

func readInvoiceFile(filename string) ([]*Invoice, error) {
	file, closer, err := openInvoiceFile(filename)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return nil, err
	}
	return readInvoices(file, suffixOf(filename))
}

func openInvoiceFile(filename string) (io.ReadCloser, func(), error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { file.Close() }
	var reader io.ReadCloser = file
	var decompressor *gzip.Reader
	if strings.HasSuffix(filename, ".gz") {
		if decompressor, err = gzip.NewReader(file); err != nil {
			return file, closer, err
		}
		closer = func() { decompressor.Close(); file.Close() }
		reader = decompressor
	}
	return reader, closer, nil
}

func readInvoices(reader io.Reader, suffix string) ([]*Invoice, error) {
	var unmarshaler InvoicesUnmarshaler
	switch suffix {
	case ".gob":
		unmarshaler = GobMarshaler{}
	case ".inv":
		unmarshaler = InvMarshaler{}
	case ".jsn", ".json":
		unmarshaler = JSONMarshaler{}
	case ".txt":
		unmarshaler = TxtMarshaler{}
	case ".xml":
		unmarshaler = XMLMarshaler{}
	}
	if unmarshaler != nil {
		return unmarshaler.UnmarshalInvoices(reader)
	}
	return nil, fmt.Errorf("unrecognized input suffix: %s", suffix)
}

func writeInvoiceFile(filename string, invoices []*Invoice) error {
	file, closer, err := createInvoiceFile(filename)
	if closer != nil {
		defer closer()
	}
	if err != nil {
		return err
	}
	return writeInvoices(file, suffixOf(filename), invoices)
}

func createInvoiceFile(filename string) (io.WriteCloser, func(), error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { file.Close() }
	var writer io.WriteCloser = file
	var compressor *gzip.Writer
	if strings.HasSuffix(filename, ".gz") {
		compressor = gzip.NewWriter(file)
		closer = func() { compressor.Close(); file.Close() }
		writer = compressor
	}
	return writer, closer, nil
}

func writeInvoices(writer io.Writer, suffix string,
	invoices []*Invoice) error {
	var marshaler InvoicesMarshaler
	switch suffix {
	case ".gob":
		marshaler = GobMarshaler{}
	case ".inv":
		marshaler = InvMarshaler{}
	case ".jsn", ".json":
		marshaler = JSONMarshaler{}
	case ".txt":
		marshaler = TxtMarshaler{}
	case ".xml":
		marshaler = XMLMarshaler{}
	}
	if marshaler != nil {
		return marshaler.MarshalInvoices(writer, invoices)
	}
	return errors.New("unrecognized output suffix")
}

func suffixOf(filename string) string {
	suffix := filepath.Ext(filename)
	if suffix == ".gz" {
		suffix = filepath.Ext(filename[:len(filename)-3])
	}
	return suffix
}

func getDepartmentId(invoiceId int) string {
	var departmentId string
	switch {
	case invoiceId < 3000:
		departmentId = "GEN"
	case invoiceId < 4000:
		departmentId = "MKT"
	case invoiceId < 5000:
		departmentId = "COM"
	case invoiceId < 6000:
		departmentId = "EXP"
	case invoiceId < 7000:
		departmentId = "INP"
	case invoiceId < 8000:
		departmentId = "TZZ"
	case invoiceId < 9000:
		departmentId = "V20"
	default:
		departmentId = "X15"
	}

	return departmentId
}

func getTaxBand(itemId string) int {
	for index, runeValue := range itemId {
		if index == 2 {
			number, _ := strconv.Atoi(string(runeValue))
			return number
		}
	}

	panic("unreachable")
}
