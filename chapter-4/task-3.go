package chapter_4

func RunTask3(data []int, columns int) [][]int {
	var newSliceLength int
	if len(data)%columns == 0 {
		newSliceLength = len(data) / columns
	} else {
		newSliceLength = len(data)/columns + 1
	}

	newSlice := make([][]int, newSliceLength)
	row, column := 0, 0

	var interimSLice []int
	for _, num := range data {
		if column == 0 {
			interimSLice = make([]int, columns)
			newSlice[row] = interimSLice
		}

		newSlice[row][column] = num

		if column == columns-1 {
			column = 0
			row += 1
		} else {
			column++
		}
	}

	return newSlice
}
