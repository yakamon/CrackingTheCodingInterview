package search

type ListWithoutSize struct {
	list []int
}

func (l *ListWithoutSize) At(i int) int {
	if len(list) <= i {
		return -1
	}
	return list[i]
}

func Search(l ListWithoutSize, x int) int {
	if l.At(0) == -1 {
		return -1
	}
	index := 1
	for l.At(index) != -1 && l.At(index) < x {
		index *= 2
	}
	return BinarySearch(l, index/2, index, x)
}

func BinarySearch(l ListWithoutSize, l, r, x int) int {
	for l <= r {
		m := (l + r) / 2
		mv := l.At(m)
		if x < mv || m == -1 {
			r = m - 1
		} else if mv < x {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
