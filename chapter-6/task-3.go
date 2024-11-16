package chapter_6

import (
	"fmt"
	"programming-in-Go-by-Mark-Summerfield/chapter-6/ordered-slice"
)

func RunTask3() {
	orderedSlice := ordered_slice.NewIntOrderedSlice([]int{1, 1, 1, 1})
	fmt.Println(orderedSlice.Index(1))

	fmt.Println(orderedSlice)
	fmt.Printf("Slice length is %d\n", orderedSlice.Len())
	orderedSlice.Add(8)
	fmt.Println(orderedSlice)
	fmt.Printf("Slice length is %d\n", orderedSlice.Len())
	orderedSlice.Add(1)
	fmt.Println(orderedSlice)
	fmt.Printf("Slice length is %d\n", orderedSlice.Len())
	fmt.Printf("1 is in slice -> %v\n", orderedSlice.Contains(1))
	fmt.Printf("5 is in slice -> %v\n", orderedSlice.Contains(5))
	orderedSlice.Add(4)
	orderedSlice.Add(8)
	orderedSlice.Add(3)
	orderedSlice.Add(6)
	orderedSlice.Add(2)
	orderedSlice.Add(0)
	fmt.Println(orderedSlice)
	fmt.Printf("Slice length is %d\n", orderedSlice.Len())
	fmt.Printf("4 is at index %d\n", orderedSlice.Index(4))
	fmt.Printf("5 is at index %d\n", orderedSlice.Index(5))
	fmt.Printf("Value at index 0 is %d\n", orderedSlice.At(orderedSlice.Index(0)))
	orderedSlice.Remove(4)
	fmt.Printf("Slice after removal of 4 -> %v\n", orderedSlice)
	orderedSlice.Clear()
	fmt.Println(orderedSlice)
	fmt.Printf("Slice length is %d\n", orderedSlice.Len())

	//orderedSlice := ordered_slice.NewStringOrderedSlice([]string{"c", "a", "b", "f"})
	//fmt.Println(orderedSlice)
	//orderedSlice.Add("d")
	//fmt.Println(orderedSlice)
}
