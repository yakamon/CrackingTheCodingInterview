package merge

// MergeSortedArrays merge two sorted arrays.
func MergeSortedArrays(a, b []int, al, bl int) {
	if al <= 1 || bl <= 1 {
		return
	}
	ai, bi := al-1, bl-1
	for i := al + bl - 1; i >= 0 && ai >= 0 && bi >= 0; i-- {
		if a[ai] < b[bi] {
			a[i] = b[bi]
			bi--
		} else {
			a[i] = a[ai]
			ai--
		}
	}
}
