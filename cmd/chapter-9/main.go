package main

import (
	"fmt"
	chapter_9 "programming-in-Go-by-Mark-Summerfield/chapter-9"
)

func main() {
	if links, err := chapter_9.RunTask1("https://www.larstornoe.com/"); err == nil {
		fmt.Printf("Links: %v\n", links)
	} else {
		fmt.Printf("Error %v\n", err)
	}
}
