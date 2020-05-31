package chapter01

// HasSameChars check if s1 has the same characters as s2
func HasSameChars(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	chars := map[rune]int{}
	for _, c := range s1 {
		if _, ok := chars[c]; ok {
			chars[c]++
			continue
		}
		chars[c] = 1
	}

	for _, c := range s2 {
		if cnt, ok := chars[c]; cnt == 0 || !ok {
			return false
		}
		chars[c]--
	}
	return true
}
