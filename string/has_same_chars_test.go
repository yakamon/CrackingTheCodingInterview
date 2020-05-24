package chapter01

import "testing"

func TestHasSameChars(t *testing.T) {
	table := []struct {
		TestData [2]string
		Expected bool
	}{
		{[2]string{"acb", "abc"}, true},
		{[2]string{"najpa", "japqn"}, false},
		{[2]string{"tfhie", "iehtf"}, true},
		{[2]string{"t", "f"}, false},
	}

	for _, data := range table {
		if actual := HasSameChars(data.TestData[0], data.TestData[1]); actual != data.Expected {
			t.Errorf("Actual: %v, Expected: %v", actual, data.Expected)
		}
	}
}
