package powerset

import (
	"strconv"
)

// PowerSet returns power set of set.
func PowerSet(set []int) [][]int {
	return powerSet(set, 0, [][]int{})
}

func powerSet(set []int, index int, memo [][]int) [][]int {
	if index == len(set) {
		return append(memo, []int{})
	}
	memo = powerSet(set, index+1, memo)
	n := set[index]
	size := len(memo)
	for i := 0; i < size; i++ {
		memo = append(memo, append([]int{n}, memo[i]...))
	}
	return memo
}

// PowerSet2 returns power set of set.
func PowerSet2(set []int) []string {
	return powerSet2(set, []string{})
}

func powerSet2(set []int, memo []string) []string {
	if len(set) == 0 {
		return append(memo, "")
	}
	n, rest := set[0], set[1:]
	memo = powerSet2(rest, memo)
	currentMemo := make([]string, len(memo), len(memo))
	for i := range currentMemo {
		currentMemo[i] = strconv.Itoa(n) + memo[i]
	}
	return append(memo, currentMemo...)
}

// PowerSet3 returns power set of set.
func PowerSet3(set []int) [][]int {
	powerSet := make([][]int, 1<<len(set))
	powerSet[0] = []int{}
	for i := 0; i < len(set); i++ {
		k := 1 << i
		for j := 0; j < k; j++ {
			powerSet[k+j] = copyAppend3(powerSet[j], set[i])
		}
	}
	return powerSet
}

// PowerSet4 returns power set of set.
func PowerSet4(set []int) [][]int {
	powerSet := make([][]int, 1<<len(set))
	powerSet[0] = []int{}
	for i, n := range set {
		k := 1 << i
		for j := 0; j < k; j++ {
			powerSet[k+j] = copyAppend(powerSet[j], n)
		}
	}
	return powerSet
}

func copyAppend(slice []int, n int) []int {
	copied := make([]int, 0, len(slice)+1)
	copied = append(copied, slice...)
	return append(copied, n)
}

func copyAppend2(nums []int, n int) []int {
	copied := make([]int, len(nums)+1, len(nums)+1)
	copy(copied, nums)
	copied[len(copied)-1] = n
	return copied
}

func copyAppend3(nums []int, n int) []int {
	return append(append(make([]int, 0, len(nums)+1), nums...), n)
}

// PowerSet5 returns power set of set.
func PowerSet5(set []int) [][]int {
	size := 1 << len(set)
	powerSet := make([][]int, size, size)
	for k := 0; k < size; k++ {
		powerSet = append(powerSet, convertIntToSlice(k, set))
	}
	return powerSet
}

func convertIntToSlice(k int, set []int) []int {
	slice := []int{}
	for i := 0; i < len(set); i++ {
		if (k & (1 << i)) != 1 {
			slice = append(slice, set[i])
		}
	}
	return slice
}
