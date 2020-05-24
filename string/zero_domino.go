package chapter01

// ZeroDomino fills rows and cols in matrix containing 0
func ZeroDomino(matrix [][]int) [][]int {
	fillRows, fillCols := map[int]bool{}, map[int]bool{}
	for rowIndex, row := range matrix {
		for colIndex, value := range row {
			if value == 0 {
				fillRows[rowIndex] = true
				fillCols[colIndex] = true
			}
		}
	}

	for rowIndex := range fillRows {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[rowIndex][i] = 0
		}
	}

	for i := 0; i < len(matrix); i++ {
		for colIndex := range fillCols {
			matrix[i][colIndex] = 0
		}
	}

	return matrix
}

// ZeroDomino2 fills rows and cols in matrix containing 0
func ZeroDomino2(matrix [][]int) [][]int {
	// Check if first row and col have zero
	isZeroFirstRow, isZeroFirstCol := false, false
	for coli := 0; coli < len(matrix[0]); coli++ {
		if matrix[0][coli] == 0 {
			isZeroFirstRow = true
		}
	}
	for rowi := 0; rowi < len(matrix); rowi++ {
		if matrix[rowi][0] == 0 {
			isZeroFirstCol = true
		}
	}

	// Search zeros
	for rowi := 1; rowi < len(matrix); rowi++ {
		for coli := 1; coli < len(matrix[0]); coli++ {
			if matrix[rowi][coli] == 0 {
				matrix[rowi][0] = 0
				matrix[0][coli] = 0
			}
		}
	}

	// Fill zero row
	for rowi := 1; rowi < len(matrix); rowi++ {
		if matrix[rowi][0] == 0 {
			matrix = fillZeroRow(matrix, rowi)
		}
	}

	// Fill zero col
	for coli := 1; coli < len(matrix[0]); coli++ {
		if matrix[0][coli] == 0 {
			matrix = fillZeroCol(matrix, coli)
		}
	}

	if isZeroFirstRow {
		matrix = fillZeroRow(matrix, 0)
	}
	if isZeroFirstCol {
		matrix = fillZeroCol(matrix, 0)
	}

	return matrix
}

func fillZeroRow(matrix [][]int, rowi int) [][]int {
	for coli := 0; coli < len(matrix[0]); coli++ {
		matrix[rowi][coli] = 0
	}

	return matrix
}

func fillZeroCol(matrix [][]int, coli int) [][]int {
	for rowi := 0; rowi < len(matrix); rowi++ {
		matrix[rowi][coli] = 0
	}

	return matrix
}
