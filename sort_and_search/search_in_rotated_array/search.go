package search

// SearchRotatedArray searchs for element x in rotated array.
func SearchRotatedArray(s []int, x int) int {
	if len(s) == 0 {
		return -1
	}
	return searchRotateArray(s, 0, len(s)-1, x)
}

func searchRotateArray(s []int, l, r, x int) int {
	m := (l + r) / 2
	if s[m] == x {
		return m
	}
	if l > r {
		return -1
	}

	if s[l] < s[m] {
		if s[l] <= x && x < s[m] {
			return searchRotateArray(s, l, m-1, x)
		}
		return searchRotateArray(s, m+1, r, x)
	} else if s[l] > s[m] {
		if s[m] < x && x <= s[r] {
			return searchRotateArray(s, m+1, r, x)
		}
		return searchRotateArray(s, l, m-1, x)
	} else if s[l] == s[m] {
		if s[m] != s[r] {
			return searchRotateArray(s, m+1, r, x)
		}
		ret := searchRotateArray(s, l, m-1, x)
		if ret != 1 {
			return searchRotateArray(s, m+1, r, x)
		}
		return ret
	}

	return -1
}
