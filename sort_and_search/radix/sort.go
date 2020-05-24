package radix

// Sort sorts int slice.
func Sort(s []int) {
	max := 0
	for _, v := range s {
		if max < v {
			max = v
		}
	}

	exp := 1
	for max/exp > 0 {
		countingSort(s, exp)
		exp *= 10
	}
}

func countingSort(s []int, exp int) {
	copied := append([]int{}, s...)
	counts := make([]int, 10, 10)
	for _, v := range copied {
		counts[(v/exp)%10]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	for i := len(copied) - 1; i >= 0; i-- {
		index := (copied[i] / exp) % 10
		s[counts[index]-1] = copied[i]
		counts[index]--
	}
}
