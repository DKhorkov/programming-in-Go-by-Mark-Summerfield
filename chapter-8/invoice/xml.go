package invoice

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
	"time"
)

type XMLMarshaler struct{}

type XMLInvoices struct {
	XMLName xml.Name      `xml:"INVOICES"`
	Version int           `xml:"version,attr"`
	Invoice []*XMLInvoice `xml:"INVOICE"`
}

type XMLInvoice struct {
	XMLName      xml.Name   `xml:"INVOICE"`
	Id           int        `xml:",attr"`
	DepartmentId string     `xml:",attr"`
	CustomerId   int        `xml:",attr"`
	Raised       string     `xml:",attr"`
	Due          string     `xml:",attr"`
	Paid         bool       `xml:",attr"`
	Note         string     `xml:"NOTE"`
	Item         []*XMLItem `xml:"ITEM"`
}

type XMLItem struct {
	XMLName  xml.Name `xml:"ITEM"`
	Id       string   `xml:",attr"`
	Price    float64  `xml:",attr"`
	Quantity int      `xml:",attr"`
	TaxBand  int      `xml:",attr"`
	Note     string   `xml:"NOTE"`
}

func XMLInvoicesForInvoices(invoices []*Invoice) *XMLInvoices {
	xmlInvoices := &XMLInvoices{
		Version: fileVersion,
		Invoice: make([]*XMLInvoice, 0, len(invoices)),
	}
	for _, invoice := range invoices {
		xmlInvoices.Invoice = append(xmlInvoices.Invoice,
			XMLInvoiceForInvoice(invoice))
	}
	return xmlInvoices
}

func XMLInvoiceForInvoice(invoice *Invoice) *XMLInvoice {
	xmlInvoice := &XMLInvoice{
		Id:           invoice.Id,
		CustomerId:   invoice.CustomerId,
		DepartmentId: invoice.DepartmentId,
		Raised:       invoice.Raised.Format(dateFormat),
		Due:          invoice.Due.Format(dateFormat),
		Paid:         invoice.Paid,
		Note:         invoice.Note,
		Item:         make([]*XMLItem, 0, len(invoice.Items)),
	}
	for _, item := range invoice.Items {
		xmlItem := &XMLItem{
			Id:       item.Id,
			Price:    item.Price,
			Quantity: item.Quantity,
			TaxBand:  item.TaxBand,
			Note:     item.Note,
		}
		xmlInvoice.Item = append(xmlInvoice.Item, xmlItem)
	}
	return xmlInvoice
}

func (xmlInvoices *XMLInvoices) Invoices() (invoices []*Invoice,
	err error) {
	invoices = make([]*Invoice, 0, len(xmlInvoices.Invoice))
	for _, xmlInvoice := range xmlInvoices.Invoice {
		invoice, err := xmlInvoice.Invoice()
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (xmlInvoice *XMLInvoice) Invoice() (invoice *Invoice, err error) {
	invoice = &Invoice{
		Id:           xmlInvoice.Id,
		CustomerId:   xmlInvoice.CustomerId,
		DepartmentId: getDepartmentId(xmlInvoice.Id),
		Paid:         xmlInvoice.Paid,
		Note:         strings.TrimSpace(xmlInvoice.Note),
		Items:        make([]*Item, 0, len(xmlInvoice.Item)),
	}
	if invoice.Raised, err = time.Parse(dateFormat, xmlInvoice.Raised); err != nil {
		return nil, err
	}
	if invoice.Due, err = time.Parse(dateFormat, xmlInvoice.Due); err != nil {
		return nil, err
	}
	for _, xmlItem := range xmlInvoice.Item {
		item := &Item{
			Id:       xmlItem.Id,
			Price:    xmlItem.Price,
			Quantity: xmlItem.Quantity,
			TaxBand:  getTaxBand(xmlItem.Id),
			Note:     strings.TrimSpace(xmlItem.Note),
		}
		invoice.Items = append(invoice.Items, item)
	}
	return invoice, nil
}

func (XMLMarshaler) MarshalInvoices(writer io.Writer,
	invoices []*Invoice) error {
	if _, err := writer.Write([]byte(xml.Header)); err != nil {
		return err
	}
	xmlInvoices := XMLInvoicesForInvoices(invoices)
	encoder := xml.NewEncoder(writer)
	return encoder.Encode(xmlInvoices)
}

func (XMLMarshaler) UnmarshalInvoices(reader io.Reader) ([]*Invoice,
	error) {
	xmlInvoices := &XMLInvoices{}
	decoder := xml.NewDecoder(reader)
	if err := decoder.Decode(xmlInvoices); err != nil {
		return nil, err
	}
	if xmlInvoices.Version > fileVersion {
		return nil, fmt.Errorf("version %d is too new to read",
			xmlInvoices.Version)
	}
	return xmlInvoices.Invoices()
}
