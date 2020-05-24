package chapter01

// IsConvertibleByOneEdit1 checks if a string is able to be converted to another string
// by only one operation of replacing, inserting or removing a character
func IsConvertibleByOneEdit1(s, t string) bool {
	longStr, shortStr, lenDiff := getLongShortStringsAndLenDiff([]rune(s), []rune(t))
	switch lenDiff {
	case 0:
		return checkConvertibleByReplace(longStr, shortStr)
	case 1:
		return checkConvertibleByRemove(longStr, shortStr)
	default:
		return false
	}
}

func getLongShortStringsAndLenDiff(s, t []rune) ([]rune, []rune, int) {
	lenDiff := len(s) - len(t)
	if lenDiff < 0 {
		return t, s, -lenDiff
	}
	return s, t, lenDiff
}

func checkConvertibleByReplace(s, t []rune) bool {
	foundMismatch := false
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			if foundMismatch {
				return false
			}
			foundMismatch = true
		}
	}
	return true
}

func checkConvertibleByRemove(longStr, shortStr []rune) bool {
	foundMismatch := false
	shortStrIndex := 0
	for longStrIndex := 0; longStrIndex < len(shortStr); longStrIndex++ {
		if longStr[longStrIndex] != shortStr[shortStrIndex] {
			if foundMismatch {
				return false
			}
			foundMismatch = true
			continue
		}
		shortStrIndex++
	}
	return true
}

// IsConvertibleByOneEdit2 checks if a string is able to be converted to another string
// by only one operation of replacing, inserting or removing a character
func IsConvertibleByOneEdit2(s, t string) bool {
	long, short, lenDiff := getLongShortStringsAndLenDiff([]rune(s), []rune(t))
	if lenDiff > 1 {
		return false
	}

	foundDiff := false
	for li, si := 0, 0; si < len(short); li++ {
		if short[si] != long[li] {
			if foundDiff {
				return false
			}
			foundDiff = true
			if len(short) == len(long) {
				si++
			}
		} else {
			si++
		}
	}
	return true
}
