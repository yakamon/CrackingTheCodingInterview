package ep

import (
	"testing"
)

func TestExistsPath(t *testing.T) {
	g := New()

	root := NewNode()
	g.AppendNode(root)

	layer1 := []*Node{NewNode(), NewNode()}
	root.AppendChildren(layer1...)
	g.AppendNode(layer1...)

	layer2 := [][]*Node{{NewNode(), NewNode()}, {NewNode(), NewNode()}}
	for i, parent := range layer1 {
		parent.AppendChildren(layer2[i]...)
		g.AppendNode(layer2[i]...)
	}

	layer3 := [][]*Node{{NewNode()}, {}, {NewNode(), NewNode()}, {}}
	for i := range layer2 {
		for j, parent := range layer2[i] {
			parent.AppendChildren(layer3[2*i+j]...)
			g.AppendNode(layer3[2*i+j]...)
		}
	}

	if !g.ExistsPath(root, layer3[0][0]) {
		t.Error("expected true")
	}
	if g.ExistsPath(layer1[1], layer3[0][0]) {
		t.Error("expected false")
	}
}
