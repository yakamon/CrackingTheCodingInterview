package palindrome

import (
	"testing"
)

var testDataList = [][]interface{}{
	{[]interface{}{1, 2, 3, 2, 1}, true},
	{[]interface{}{1, 2, 3, 3, 2}, false},
	{[]interface{}{1, 2, 3, 3, 2, 1}, true},
	{[]interface{}{1, 2, 3, 4, 2, 1}, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, td := range testDataList {
		node := NewBySlice(td[0].([]interface{}))
		expected := td[1].(bool)
		if actual := IsPalindrome(node); actual != expected {
			t.Errorf("The answer is incorrect.\nExpected: %v\nActual:   %v", expected, actual)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, td := range testDataList {
			node := NewBySlice(td[0].([]interface{}))
			b.Log(IsPalindrome(node))
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, td := range testDataList {
		node := NewBySlice(td[0].([]interface{}))
		expected := td[1].(bool)
		if actual := IsPalindrome2(node); actual != expected {
			t.Errorf("The answer is incorrect.\nExpected: %v\nActual:   %v", expected, actual)
		}
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, td := range testDataList {
			node := NewBySlice(td[0].([]interface{}))
			b.Log(IsPalindrome2(node))
		}
	}
}

func TestIsPalindrome3(t *testing.T) {
	for _, td := range testDataList {
		node := NewBySlice(td[0].([]interface{}))
		expected := td[1].(bool)
		if actual := IsPalindrome3(node); actual != expected {
			t.Errorf("The answer is incorrect.\nExpected: %v\nActual:   %v", expected, actual)
		}
	}
}

func BenchmarkIsPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, td := range testDataList {
			node := NewBySlice(td[0].([]interface{}))
			b.Log(IsPalindrome3(node))
		}
	}
}

func TestIsPalindrome4(t *testing.T) {
	for _, td := range testDataList {
		node := NewBySlice(td[0].([]interface{}))
		expected := td[1].(bool)
		if actual := IsPalindrome4(node); actual != expected {
			t.Errorf("The answer is incorrect.\nExpected: %v\nActual:   %v", expected, actual)
		}
	}
}

func BenchmarkIsPalindrome4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, td := range testDataList {
			node := NewBySlice(td[0].([]interface{}))
			b.Log(IsPalindrome4(node))
		}
	}
}
