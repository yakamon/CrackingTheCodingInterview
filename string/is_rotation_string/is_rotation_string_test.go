package chapter01

import "testing"

func TestIsRotationString(t *testing.T) {
	testCases := []struct {
		A   string
		B   string
		Exp bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"waterbott", "erbottlewat", false},
		{"waterbottli", "erbottlewat", false},
	}

	for _, tc := range testCases {
		if act := IsRotationString(tc.A, tc.B); act != tc.Exp {
			t.Errorf("Expected: %v, Actual: %v", tc.Exp, act)
		}
	}
}

func BenchmarkIsRotationString(b *testing.B) {
	testCases := []struct {
		A   string
		B   string
		Exp bool
	}{
		{"waterbottle", "erbottlewat", true},
		{"waterbott", "erbottlewat", false},
		{"waterbottli", "erbottlewat", false},
	}

	for iter := 0; iter < b.N; iter++ {
		for _, tc := range testCases {
			b.Log(IsRotationString(tc.A, tc.B))
		}
	}
}
