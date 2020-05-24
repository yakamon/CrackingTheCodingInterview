package selection

// Sort sorts int slice.
func Sort(s []int) {
	length := len(s)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i; j < length; j++ {
			if s[minIndex] > s[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			s[i], s[minIndex] = s[minIndex], s[i]
		}
	}
}
