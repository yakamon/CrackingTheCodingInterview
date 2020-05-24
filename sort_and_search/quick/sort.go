package quick

// Sort sorts int slice.
func Sort(s []int) {
	sortInner(s, 0, len(s)-1)
}

func sortInner(s []int, l, r int) {
	if l >= r {
		return
	}
	i := partition(s, l, r)
	if l < i-1 {
		sortInner(s, l, i-1)
	}
	if i < r {
		sortInner(s, i, r)
	}
}

func partition(s []int, l, r int) int {
	pivot := s[(l+r)/2]
	for l <= r {
		for s[l] < pivot {
			l++
		}
		for pivot < s[r] {
			r--
		}
		if l <= r {
			s[l], s[r] = s[r], s[l]
			l++
			r--
		}
	}
	return l
}
