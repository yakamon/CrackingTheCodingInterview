package bubble

// Sort sorts int slice.
func Sort(s []int) {
	length := len(s)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
