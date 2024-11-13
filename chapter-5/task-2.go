package chapter_5

import (
	"fmt"
	"unicode/utf8"
)

func RunTask2(words []string, ascii bool) {
	var isPalindromeFunc func(string) bool = utf8PalindromeFunc
	if ascii {
		isPalindromeFunc = asciiPalindromeFunc
	}

	for _, word := range words {
		fmt.Printf("%5t %q\n", isPalindromeFunc(word), word)
	}
}

func utf8PalindromeFunc(s string) bool { // UTF-8 version
	//if utf8.RuneCountInString(s) <= 1 {
	//	return true
	//}
	//
	//first, sizeOfFirst := utf8.DecodeRuneInString(s)
	//last, sizeOfLast := utf8.DecodeLastRuneInString(s)
	//if first != last {
	//	return false
	//}
	//return utf8PalindromeFunc(s[sizeOfFirst : len(s)-sizeOfLast])

	for utf8.RuneCountInString(s) > 1 {
		first, sizeOfFirst := utf8.DecodeRuneInString(s)
		last, sizeOfLast := utf8.DecodeLastRuneInString(s)
		if first != last {
			return false
		}

		s = s[sizeOfFirst : len(s)-sizeOfLast]
	}

	return true
}

// asciiPalindromeFunc is a simple ASCII-only version
func asciiPalindromeFunc(s string) bool {
	//if len(s) <= 1 {
	//	return true
	//}
	//if s[0] != s[len(s)-1] {
	//	return false
	//}
	//
	//return asciiPalindromeFunc(s[1 : len(s)-1])

	for len(s) > 1 {
		if s[0] != s[len(s)-1] {
			return false
		}

		s = s[1 : len(s)-1]
	}

	return true
}
