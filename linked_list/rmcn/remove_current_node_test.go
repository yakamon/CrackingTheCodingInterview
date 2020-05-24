package rmcn

import (
	"testing"
)

func TestRemoveCurrentNode(t *testing.T) {
	testData := []interface{}{1, 3, 5, 7, 9, 11}
	node := NewNode(testData[0])
	node.AppendListToTail(testData[1:])

	delIndex := 2
	delValue := testData[delIndex]
	delNode := node.Next.Next
	RemoveCurrentNode(delNode)
	if node.Contains(delValue) {
		t.Errorf("Failed to remove %vth value %v", delIndex, delValue)
	}

	delIndex = 1
	delValue = testData[delIndex]
	delNode = node.Next
	RemoveCurrentNode(delNode)
	if node.Contains(delValue) {
		t.Errorf("Failed to remove %vth value %v", delIndex, delValue)
	}
}
