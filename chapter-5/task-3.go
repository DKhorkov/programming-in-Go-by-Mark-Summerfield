package chapter_5

import (
	"bytes"
	"sort"
)

func RunTask3(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	if len(words) == 0 || len(words[0]) == 0 {
		return ""
	}

	var commonPrefix bytes.Buffer

	runeWords := make([][]rune, 0, len(words))
	for _, word := range words {
		runeWords = append(runeWords, []rune(word))
	}

BREAKPOINT:
	for charIndex := 0; charIndex < len(runeWords[0]); charIndex++ {
		char := runeWords[0][charIndex]
		for nextWordIndex := 1; nextWordIndex < len(runeWords); nextWordIndex++ {
			if char != runeWords[nextWordIndex][charIndex] {
				break BREAKPOINT
			}
		}

		commonPrefix.WriteRune(char)
	}

	return commonPrefix.String()
}
