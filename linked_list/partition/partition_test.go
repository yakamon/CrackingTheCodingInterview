package partition

import (
	"fmt"
	"testing"
)

func TestByPivot1(t *testing.T) {
	testData := []interface{}{3, 5, 8, 5, 10, 2, 1}
	node := NewNode(testData[0])
	node.AppendListToTail(testData[1:])
	fmt.Println(node.Values())

	node = ByPivot1(node, 5)
	fmt.Println(node.Values())
}

func TestByPivot2(t *testing.T) {
	testData := []interface{}{3, 5, 8, 5, 10, 2, 1}
	node := NewNode(testData[0])
	node.AppendListToTail(testData[1:])
	fmt.Println(node.Values())

	node = ByPivot2(node, 5)
	fmt.Println(node.Values())
}
