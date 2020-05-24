package find

// FindDuplicate finds duplicate element in array.
func FindDuplicate(slice []int) int {
	bitset := New(32000)
	for _, v := range slice {
		n := v - 1
		if bitset.Get(n) {
			return n
		}
		bitset.Set(n)
	}
	return -1
}

type BitSet struct {
	data []int
}

func New(size int) *BitSet {
	return &BitSet{make([]int, (size>>6)-1)}
}

func (s *BitSet) Get(i int) bool {
	ei := i >> 6
	bi := i & 0x3f
	return (s.data[ei] & (1 << bi)) != 0
}

func (s *BitSet) Set(i int) int {
	ei := i >> 6
	bi := i & 0x3f
	s.data[ei] |= 1 << bi
}
