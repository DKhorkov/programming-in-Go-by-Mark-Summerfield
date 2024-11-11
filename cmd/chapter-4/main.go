package main

import (
	"fmt"
	"programming-in-Go-by-Mark-Summerfield/chapter-4"
)

func main() {
	// Task 1:
	nonUnique := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	fmt.Printf(
		"Non unique %v -> unique %v\n",
		nonUnique,
		chapter_4.RunTask1(nonUnique),
	)

	// Task 2:
	twoDimensionalSlice := [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11},
		[]int{12, 13, 14, 15},
		[]int{16, 17, 18, 19, 20},
	}
	fmt.Printf(
		"Two dimensional slice %v -> single dimensional slice %v\n",
		twoDimensionalSlice,
		chapter_4.RunTask2(twoDimensionalSlice),
	)

	// Task 3:
	task3Data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	columns := 3
	fmt.Printf(
		"Array %v to two dimensional array with %d columns -> %v",
		task3Data,
		columns,
		chapter_4.RunTask3(task3Data, columns),
	)

	// Task 4:
	task4Data := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}

	parsedIni := chapter_4.RunTask4(task4Data)
	fmt.Printf(
		"%v mapped by groups -> %v\n",
		task4Data,
		parsedIni,
	)

	// Task 5:
	chapter_4.RunTask5(parsedIni)
}
