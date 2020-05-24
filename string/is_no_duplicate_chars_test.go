package chapter01

import "testing"

func TestIsNonDuplicateString(t *testing.T) {
	table := []struct {
		TestData string
		Expected bool
	}{
		{"acb", true},
		{"najpa", false},
		{"tfhieiehtf", false},
		{"tf", true},
	}

	for _, data := range table {
		if actual := IsNonDuplicateString(data.TestData); actual != data.Expected {
			t.Errorf("Actual: %v, Expected: %v", actual, data.Expected)
		}
	}
}
