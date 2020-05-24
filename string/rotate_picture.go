package chapter01

// RotatePicture rotate matrix-defined picture right
func RotatePicture(pic [][]int) [][]int {
	sideLen := len(pic)
	newPic := make([][]int, sideLen)
	for x := 0; x <= sideLen-1; x++ {
		newRow := make([]int, sideLen)
		for y, i := sideLen-1, 0; y >= 0; y, i = y-1, i+1 {
			newRow[i] = pic[y][x]
		}
		newPic[x] = newRow
	}
	return newPic
}

// RotatePicture2 rotate matrix-defined picture
func RotatePicture2(pic [][]int) [][]int {
	if len(pic) == 0 || len(pic) != len(pic[0]) {
		return nil
	}

	n := len(pic)
	for layer := 0; layer < n/2; layer++ {
		// first and last
		f, l := layer, n-1-layer
		for i := f; i < l; i++ {
			offset := i - f
			pic[f][i], pic[l-offset][f], pic[l][l-offset], pic[i][l] = pic[l-offset][f], pic[l][l-offset], pic[i][l], pic[f][i]
		}
	}

	return pic
}
