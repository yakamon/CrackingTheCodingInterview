package chapter01

import (
	"testing"
)

var tbl = []struct {
	Case     struct{ s string }
	Expected string
}{
	{struct{ s string }{"aabcccccaaa"}, "a2b1c5a3"},
	{struct{ s string }{"abcccc"}, "abcccc"},
	{struct{ s string }{"aaaaaaaaaabbbbbbbbbbbbbbb"}, "a10b15"},
}

func TestBasicRunLengthEncode1(t *testing.T) {
	for i, test := range tbl {
		t.Logf("Test case %d", i)
		if actual := BasicRunLengthEncode1(test.Case.s); actual != test.Expected {
			t.Errorf("  Expected: %v, Got: %v", test.Expected, actual)
		}
	}
}

func BenchmarkBasicRunLengthEncode1(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, test := range tbl {
			b.Log(BasicRunLengthEncode1(test.Case.s))
		}
	}
}

func TestBasicRunLengthEncode2(t *testing.T) {
	for i, test := range tbl {
		t.Logf("Test case %d", i)
		if actual := BasicRunLengthEncode2(test.Case.s); actual != test.Expected {
			t.Errorf("  Expected: %v, Got: %v", test.Expected, actual)
		}
	}
}

func BenchmarkBasicRunLengthEncode2(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, test := range tbl {
			b.Log(BasicRunLengthEncode2(test.Case.s))
		}
	}
}

func TestBasicRunLengthEncode3(t *testing.T) {
	for i, test := range tbl {
		t.Logf("Test case %d", i)
		if actual := BasicRunLengthEncode3(test.Case.s); actual != test.Expected {
			t.Errorf("  Expected: %v, Got: %v", test.Expected, actual)
		}
	}
}

func BenchmarkBasicRunLengthEncode3(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, test := range tbl {
			b.Log(BasicRunLengthEncode3(test.Case.s))
		}
	}
}

func TestBasicRunLengthEncode4(t *testing.T) {
	for i, test := range tbl {
		t.Logf("Test case %d", i)
		if actual := BasicRunLengthEncode4(test.Case.s); actual != test.Expected {
			t.Errorf("  Expected: %v, Got: %v", test.Expected, actual)
		}
	}
}

func BenchmarkBasicRunLengthEncode4(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, test := range tbl {
			b.Log(BasicRunLengthEncode4(test.Case.s))
		}
	}
}

func TestBasicRunLengthEncode5(t *testing.T) {
	for i, test := range tbl {
		t.Logf("Test case %d", i)
		if actual := BasicRunLengthEncode5(test.Case.s); actual != test.Expected {
			t.Errorf("  Expected: %v, Got: %v", test.Expected, actual)
		}
	}
}

func BenchmarkBasicRunLengthEncode5(b *testing.B) {
	for iter := 0; iter < b.N; iter++ {
		for _, test := range tbl {
			b.Log(BasicRunLengthEncode5(test.Case.s))
		}
	}
}

func BenchmarkStringConcatWithRuneSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var rs []rune
		for _, c := range "abcdefghijklmnopqrstuvwxyz0123456789" {
			rs = append(rs, c)
		}
		s := string(rs)
		b.Log(s)
	}
}

func BenchmarkStringConcatWithStringOperator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s string
		for _, c := range "abcdefghijklmnopqrstuvwxyz0123456789" {
			s += string(c)
		}
		b.Log(s)
	}
}
