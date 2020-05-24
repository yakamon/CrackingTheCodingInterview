package chapter01

import "testing"

func TestIsConveribleByOneEdit1(t *testing.T) {
	tbl := []struct {
		s1  string
		s2  string
		exp bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"pale", "pa", false},
	}

	for i, d := range tbl {
		t.Logf("TestCase %d:", i)
		if act := IsConvertibleByOneEdit1(d.s1, d.s2); act != d.exp {
			t.Errorf("Expected: %v, but Got: %v", d.exp, act)
		}
	}
}

func TestIsConveribleByOneEdit2(t *testing.T) {
	tbl := []struct {
		s1  string
		s2  string
		exp bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"pale", "pa", false},
	}

	for i, d := range tbl {
		t.Logf("TestCase %d:", i)
		if act := IsConvertibleByOneEdit2(d.s1, d.s2); act != d.exp {
			t.Errorf("Expected: %v, but Got: %v", d.exp, act)
		}
	}
}
