package chapter_6

import (
	"fmt"
	font2 "programming-in-Go-by-Mark-Summerfield/chapter-6/font"
)

func RunTask1() {
	font := font2.New("Times-New-Roman", 20)
	fmt.Println(font)
	font.SetSize(45)
	fmt.Println(font)
	font.SetFamily("Serif")
	fmt.Println(font)
}
