package search

func FindElement1(matrix [][]int, x int) ([2]int, bool) {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && 0 <= col {
		if matrix[row][col] == x {
			return [2]int{row, col}, true
		} else if x < matrix[row][col] {
			col--
		} else {
			row++
		}
	}
	return nil, false
}
