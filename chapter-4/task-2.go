package chapter_4

func RunTask2(data [][]int) []int {
	newSlice := make([]int, 0, len(data)*3)
	for _, subSlice := range data {
		for _, v := range subSlice {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}
