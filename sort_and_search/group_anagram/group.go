package anagram

// GroupAnagrams sorts string array so that anagram strings are next to each other.
func GroupAnagrams(s []string) {
	m := map[string][]string{}

	for _, str := range s {
		k := sortString(str)
		m[k] = append(m[k], str)
	}

	i := 0
	for _, group := range m {
		for _, str := range group {
			s[i] = str
			i++
		}
	}
}

// 文字列をソートする
// ヒープソート
func sortString(s string) string {
	chars := []rune(s)
	heapSort(chars)
	return string(chars)
}

func heapSort(s []rune) {
	if len(s) <= 1 {
		return
	}
	buildHeap(s)
	for last := len(s) - 1; last > 0; last-- {
		s[0], s[last] = s[last], s[0]
		maxHeapify(s, 0, last-1)
	}
}

func buildHeap(s []rune) {
	last := len(s) - 1
	for root := (last + 1) / 2; root >= 0; root-- {
		maxHeapify(s, root, last)
	}
}

func maxHeapify(s []rune, root, last int) {
	v := s[root] // root value

	i := 2*root + 1 // get left child index
	if i < last && s[i] < s[i+1] {
		i++ // move to right child index
	}

	for i <= last && v < s[i] {
		s[root] = s[i]
		root = i

		i = 2*root + 1
		if i < last && s[i] < s[i+1] {
			i++
		}
	}

	s[root] = v
}

func minHeapify(s []rune, root, last int) {
	v := s[root]

	i := 2*root + 1
	if i < last && s[i] > s[i+1] {
		i++
	}

	for i <= last && v > s[i] {
		s[root] = s[i]
		root = i

		i = 2*root + 1
		if i < last && s[i] > s[i+1] {
			i++
		}
	}

	s[root] = v
}
