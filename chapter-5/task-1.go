package chapter_5

import (
	"archive/tar"
	"archive/zip"
	"compress/bzip2"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var FunctionForSuffix = map[string]func(string) ([]string, error){
	".gz":      GzipFileList,
	".tar":     TarFileList,
	".tar.gz":  TarFileList,
	".tgz":     TarFileList,
	".zip":     ZipFileList,
	".tar.bz2": TarFileList,
}

func RunTask1(files []string) {
	archiveFileList := ArchiveFileList

	for _, filename := range files {
		fmt.Print(filename)
		lines, err := archiveFileList(filename)
		if err != nil {
			fmt.Println(" ERROR:", err)
		} else {
			fmt.Println()
			for _, line := range lines {
				fmt.Println(" ", line)
			}
		}
	}
}

func ArchiveFileList(file string) ([]string, error) {
	if function, ok := FunctionForSuffix[Suffix(file)]; ok {
		return function(file)
	}
	return nil, errors.New("unrecognized archive")
}

func Suffix(file string) string {
	file = strings.ToLower(filepath.Base(file))
	if i := strings.LastIndex(file, "."); i > -1 {
		if file[i:] == ".bz2" || file[i:] == ".gz" || file[i:] == ".xz" {
			if j := strings.LastIndex(file[:i], "."); j > -1 && strings.HasPrefix(file[j:], ".tar") {
				return file[j:]
			}
		}
		return file[i:]
	}
	return file
}

func ZipFileList(filename string) ([]string, error) {
	zipReader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer zipReader.Close()
	var files []string
	for _, file := range zipReader.File {
		files = append(files, file.Name)
	}
	return files, nil
}

func GzipFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return []string{gzipReader.Header.Name}, nil
}

func TarFileList(filename string) ([]string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var tarReader *tar.Reader
	if strings.HasSuffix(filename, ".gz") ||
		strings.HasSuffix(filename, ".tgz") {
		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			return nil, err
		}
		tarReader = tar.NewReader(gzipReader)
	} else if strings.HasSuffix(filename, ".bz2") {
		bz2Reader := bzip2.NewReader(reader)
		tarReader = tar.NewReader(bz2Reader)
	} else {
		tarReader = tar.NewReader(reader)
	}
	var files []string
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return files, err
		}
		if header == nil {
			break
		}
		files = append(files, header.Name)
	}
	return files, nil
}
