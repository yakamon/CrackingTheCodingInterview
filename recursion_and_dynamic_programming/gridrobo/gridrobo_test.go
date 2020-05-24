package gridrobo

import "testing"

var testcase = [][][]bool{
	{
		{true, true, true, true},
		{true, true, true, true},
		{true, true, true, true},
		{true, false, true, true},
		{true, true, true, true},
	},
	{
		{true, true, true, true},
		{true, true, true, false},
		{true, false, true, true},
		{true, true, true, false},
		{true, true, false, true},
	},
}

func TestFindPath(t *testing.T) {
	for _, c := range testcase {
		t.Log(FindPath(c))
	}
}

func TestFindPath2(t *testing.T) {
	for _, c := range testcase {
		t.Log(FindPath2(c))
	}
}
