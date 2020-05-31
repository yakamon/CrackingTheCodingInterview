package chapter01

import "testing"

func TestIsPalindromePermutation1(t *testing.T) {
	tbl := []struct {
		s   string
		exp bool
	}{
		{"Tact Coa", true},
		{"Tvea iha", false},
		{"Hag & gea Benn & h", true},
	}

	for i, d := range tbl {
		if act := IsPalindromePermutation1(d.s); act != d.exp {
			t.Errorf("TestCase.%d -> Expected: %v, Got: %v", i, d.exp, act)
		}
	}
}

func TestIsPalindromePermutation2(t *testing.T) {
	tbl := []struct {
		s   string
		exp bool
	}{
		{"Tact Coa", true},
		{"Tvea iha", false},
		{"Hag & gea Benn & h", true},
	}

	for i, d := range tbl {
		t.Logf("TestCase %d\n", i)
		if act := IsPalindromePermutation2(d.s); act != d.exp {
			t.Errorf("    Expected: %v, Got: %v", d.exp, act)
		}
	}
}
