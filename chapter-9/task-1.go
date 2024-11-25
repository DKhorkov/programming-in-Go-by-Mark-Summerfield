package chapter_9

import "programming-in-Go-by-Mark-Summerfield/chapter-9/linkutil"

func RunTask1(url string) ([]string, error) {
	return linkutil.LinksFromUrl(url)
}
