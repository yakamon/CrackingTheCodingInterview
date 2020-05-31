package chapter01

// ConvertSpaces convert blank spaces in string to "%20"
func ConvertSpaces(s string, trueSize int) string {
	// Count spaces to know new length of string
	spaceCnt := 0
	for _, c := range s {
		if c == ' ' {
			spaceCnt++
		}
	}

	// If space is not found, return original string
	if spaceCnt == 0 {
		return s
	}

	newSize := trueSize + spaceCnt*2
	newStr := make([]rune, newSize)
	index := newSize - 1
	for i := trueSize - 1; i >= 0; i-- {
		if rune(s[i]) == ' ' {
			newStr[index], newStr[index-1], newStr[index-2] = '0', '2', '%'
			index -= 3
		} else {
			newStr[index] = rune(s[i])
			index--
		}
	}

	return string(newStr)
}
