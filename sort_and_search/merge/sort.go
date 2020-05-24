package merge

// Sort sorts int slice.
func Sort(s []int) {
	h := make([]int, len(s), len(s))
	sortInner(s, h, 0, len(s)-1)
}

func sortInner(s, h []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	sortInner(s, h, l, m)
	sortInner(s, h, m+1, r)
	merge(s, h, l, m, r)
}

func merge(s, h []int, l, m, r int) {
	copy(h[l:r+1], s[l:r+1])

	li, ri := l, m+1
	cur := l
	for li <= m && ri <= r {
		if h[li] < h[ri] {
			s[cur] = h[li]
			li++
		} else {
			s[cur] = h[ri]
			ri++
		}
		cur++
	}

	if li <= m {
		copy(s[cur:], h[li:m+1])
	} else if ri <= r {
		copy(s[cur:], h[ri:r+1])
	}
}
