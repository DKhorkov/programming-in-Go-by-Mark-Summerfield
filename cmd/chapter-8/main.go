package main

import (
	"programming-in-Go-by-Mark-Summerfield/chapter-8"
)

func main() {
	chapter_8.RunTask1("./cmd/chapter-8/test.tar.bz2")
	for _, suffix := range []string{".json", ".xml", ".txt", ".inv"} {
		chapter_8.RunTask3(
			"./cmd/chapter-8/invoices.gob",
			"./cmd/chapter-8/invoices"+suffix,
			true,
		)
	}

	chapter_8.RunTask3(
		"./cmd/chapter-8/invoices.inv",
		"./cmd/chapter-8/invoices2.json",
		true,
	)
}
