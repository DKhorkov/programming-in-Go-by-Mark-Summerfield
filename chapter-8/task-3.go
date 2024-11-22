package chapter_8

import (
	"programming-in-Go-by-Mark-Summerfield/chapter-8/invoice"
)

func RunTask3(inFilename, outFilename string, report bool) {
	invoice.ProcessInvoices(inFilename, outFilename, report)
}
