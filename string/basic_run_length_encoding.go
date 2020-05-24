package chapter01

import (
	"fmt"
	"strconv"
)

// BasicRunLengthEncode1 encode string by Run Length Encoding
func BasicRunLengthEncode1(s string) string {
	var encoded []rune
	runeSlice := []rune(s)
	char := runeSlice[0]
	charCnt := 1
	for _, c := range runeSlice[1:] {
		if char != c {
			encoded = appendCharAndNumber(encoded, char, charCnt)
			char = c
			charCnt = 1
		} else {
			charCnt++
		}
	}
	encoded = appendCharAndNumber(encoded, char, charCnt)

	if len(runeSlice) <= len(encoded) {
		return s
	}
	return string(encoded)
}

// BasicRunLengthEncode2 encode string by Run Length Encoding
func BasicRunLengthEncode2(s string) string {
	var encoded string
	runeSlice := []rune(s)
	char := runeSlice[0]
	charCnt := 1
	for _, c := range runeSlice[1:] {
		if char != c {
			encoded += string(char) + strconv.Itoa(charCnt)
			char = c
			charCnt = 1
		} else {
			charCnt++
		}
	}
	encoded += string(char) + strconv.Itoa(charCnt)

	if len(runeSlice) <= len(encoded) {
		return s
	}
	return encoded
}

func appendCharAndNumber(encoded []rune, char rune, charCnt int) []rune {
	encoded = append(encoded, char)
	return append(encoded, []rune(strconv.Itoa(charCnt))...)
}

// BasicRunLengthEncode3 encode string with RLE
func BasicRunLengthEncode3(s string) string {
	var encoded string
	runeSlice := []rune(s)
	char := runeSlice[0]
	charCnt := 1
	for _, c := range runeSlice[1:] {
		if char != c {
			encoded += fmt.Sprintf("%c%v", char, charCnt)
			char = c
			charCnt = 1
		} else {
			charCnt++
		}
	}
	encoded += fmt.Sprintf("%c%v", char, charCnt)

	if len(s) <= len(encoded) {
		return s
	}
	return encoded
}

// BasicRunLengthEncode4 encode string with RLE
// Best Performance
func BasicRunLengthEncode4(s string) string {
	var encoded string
	runeSlice := []rune(s)
	char, startIndex := runeSlice[0], 0
	for i, c := range runeSlice[1:] {
		if char != c {
			encoded += string(char) + strconv.Itoa(i-startIndex+1)
			char = c
			startIndex = i + 1
		}
	}
	encoded += string(char) + strconv.Itoa(len(s)-startIndex)

	if len(s) <= len(encoded) {
		return s
	}
	return encoded
}

// BasicRunLengthEncode5 encode string with RLE
func BasicRunLengthEncode5(s string) string {
	var runes []rune
	var counts []int

	runeSlice := []rune(s)
	char, startIndex := runeSlice[0], 0
	for i, c := range runeSlice[1:] {
		if char != c {
			runes = append(runes, char)
			counts = append(counts, i-startIndex+1)
			char = c
			startIndex = i + 1
		}
	}
	runes = append(runes, char)
	counts = append(counts, len(s)-startIndex)

	var encoded string
	for i := 0; i < len(runes); i++ {
		encoded += string(runes[i]) + strconv.Itoa(counts[i])
	}

	if len(s) <= len(encoded) {
		return s
	}
	return encoded
}
