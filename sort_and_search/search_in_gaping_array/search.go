package search

import "strings"

func SearchInGapingArray(strs []string, str string) int {
	if len(strs) == 0 || len(str) == 0 {
		return -1
	}
	return searchInGapingArray(strs, str, 0, len(slice)-1)
}

func searchInGapingArray(strs []string, str string, first, last int) int {
	if first > last {
		return -1
	}

	m := (first + last) / 2
	for {
		if len(strs[m] == 0) {
			l := m - 1
			r := m + 1
			if l < first || last < r {
				return -1
			} else if r <= last && len(strs[r]) != 0 {
				m = r
				break
			} else if first <= l && len(strs[l]) != 0 {
				m = l
				break
			}
			l--
			r++
		}
	}

	if cmp := strings.Compare(strs[m], str); cmp == 0 {
		return m
	} else if cmp < 0 {
		return searchInGapingArray(strs, str, m+1, last)
	} else {
		return searchInGapingArray(strs, str, first, m-1)
	}
}

func SearchInGapingArray2(strs []string, str string) int {
	if len(strs) == 0 || len(str) == 0 {
		return -1
	}

	first, last := 0, len(strs)-1
	for l <= r {
		m = (first + last) / 2
		for {
			if len(strs[m]) == 0 {
				l := m - 1
				r := m + 1
				if l < first || last < r {
					return -1
				} else if r <= last && len(strs[r]) != 0 {
					m = r
					break
				} else if first <= l && len(strs[l]) != 0 {
					m = l
					break
				}
				l--
				r++
			}
		}
		if cmp := strings.Compare(strs[m], str); cmp == 0 {
			return m
		} else if cmp < 0 {
			first = m + 1
		} else {
			last = m - 1
		}
	}

	return -1
}
