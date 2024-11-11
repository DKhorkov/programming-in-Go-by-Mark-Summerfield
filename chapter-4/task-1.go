package chapter_4

func RunTask1(data []int) []int {
	newSlice := make([]int, 0, len(data)/2)
	seen := make(map[int]bool, len(data)/2)
	for _, v := range data {
		if _, found := seen[v]; !found {
			seen[v] = true
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}
