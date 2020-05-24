package search

// BinarySearch search x in s.
func BinarySearch(s []int, x int) int {
	l, r := 0, len(s)-1
	for l <= r {
		m := (l + r) / 2
		if s[m] < x {
			l = m + 1
		} else if x < s[m] {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
