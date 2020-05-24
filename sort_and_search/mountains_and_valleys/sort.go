package sort

func SortMountainsAndValleysAfterSort(slice []int) {
	mergeSort(slice)
	for i := 1; i < len(slice); i += 2 {
		slice[i-1], slice[i] = slice[i], slice[i-1]
	}
}

func mergeSort(slice []int) {
	if len(slice) <= 1 {
		return
	}
	copied := make([]int, len(slice), len(slice))
	return mergeSortInner(slice, copied, 0, len(slice)-1)
}

func mergeSortInner(slice, copied []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	mergeSortInner(slice, l, m)
	mergeSortInner(slice, m+1, r)
	merge(slice, l, m, r)
}

func merge(slice, copied []int, l, m, r int) {
	copy(copied[l:r+1], slice[l:r+1])

	i, li, ri := l, l, m+1
	for li <= m && ri <= r {
		if copied[li] < copied[ri] {
			slice[i] = copied[li]
			li++
		} else {
			slice[i] = copied[ri]
			ri++
		}
		i++
	}

	if li <= m {
		copy(slice[i:], copied[li:m+1])
	} else if ri <= m {
		copy(slice[i:], copied[ri:r+1])
	}
}

func SortMountainsAndValleysAfterSort2(slice []int) {
	for i := 1; i < len(slice); i += 2 {
		maxIndex := i - 1
		max := slice[maxIndex]
		if i < len(slice) && max < slice[i] {
			maxIndex = i
			max = slice[maxIndex]
		}
		if i+1 < len(slice) && max < slice[i+1] {
			maxIndex = i + 1
			max = slice[maxIndex]
		}

		if i != maxIndex {
			slice[i], slice[maxIndex] = slice[maxIndex], slice[i]
		}
	}
}
