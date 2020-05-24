package counting

const (
	// MaxUInt8 is max uint8.
	MaxUInt8 = ^uint8(0)
)

// Sort sorts int slice.
func Sort(s []int) {
	copied := append([]int{}, s...)
	counts := make([]int, MaxUInt8, MaxUInt8)
	for _, v := range copied {
		counts[v]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	for _, v := range copied {
		s[counts[v]-1] = v
		counts[v]--
	}
}
