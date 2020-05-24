package chapter01

import (
	"unicode"
)

// IsPalindromePermutation1 check if a string is a palindrome permutation
// 空白文字は無視、その他の記号は無視しない
func IsPalindromePermutation1(s string) bool {
	charSet := map[rune]int{}
	for _, c := range s {
		if c != ' ' {
			c = unicode.ToLower(c)
			if _, exists := charSet[c]; !exists {
				charSet[c] = 1
				continue
			}
			charSet[c]++
		}
	}

	oddFound := false
	for _, cnt := range charSet {
		if cnt%2 == 1 {
			if !oddFound {
				oddFound = true
				continue
			}
			return false
		}
	}

	return true
}

// IsPalindromePermutation2 check if a string is a palindrome permutation
// 記号はすべて無視
func IsPalindromePermutation2(s string) bool {
	bitVector := createBitVector(s)
	return bitVector == 0 || checkExactlyOneBitSet(bitVector)
}

func createBitVector(s string) (bitVector int) {
	for _, c := range s {
		bitVector = toggle(bitVector, getCharIndex(unicode.ToLower(c)))
	}
	return
}

func getCharIndex(c rune) int {
	a, z, v := int('a'), int('z'), int(c)
	if a <= v && v <= z {
		return v - a
	}
	return -1
}

func toggle(bitVector, index int) int {
	if index < 0 {
		return bitVector
	}

	mask := 1 << index
	if (bitVector & mask) == 0 {
		bitVector |= mask
	} else {
		bitVector &= ^mask
	}

	return bitVector
}

func checkExactlyOneBitSet(bitVector int) bool {
	return (bitVector & (bitVector - 1)) == 0
}
