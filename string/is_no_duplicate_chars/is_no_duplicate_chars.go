package chapter01

// IsNonDuplicateString checks if a string contains no duplicate charactors
func IsNonDuplicateString(s string) bool {
	chars := map[rune]bool{}
	for _, c := range s {
		if _, ok := chars[c]; !ok {
			chars[c] = true
			continue
		}
		return false
	}
	return true
}
