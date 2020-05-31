package chapter01

import (
	"testing"
)

func TestConvertSpaces(t *testing.T) {
	table := []struct {
		s        string
		n        int
		expected string
	}{
		{"a b c", 5, "a%20b%20c"},
	}

	for _, d := range table {
		if actual := ConvertSpaces(d.s, d.n); actual != d.expected {
			t.Errorf("Expected: %v, Actual: %v", d.expected, actual)
		}
	}
}
