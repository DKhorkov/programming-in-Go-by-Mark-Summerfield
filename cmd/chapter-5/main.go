package main

import (
	"fmt"
	"programming-in-Go-by-Mark-Summerfield/chapter-5"
)

func main() {
	// Task 2:
	words := []string{"dollar", "level"}
	chapter_5.RunTask2(words, false)
	chapter_5.RunTask2(words, true)

	// Task 3:
	commonPrefixWords := []string{"commonUncommon", "commonPrefix", "co"}
	fmt.Printf("common prefix for %v is %s\n", commonPrefixWords, chapter_5.RunTask3(commonPrefixWords))
}
