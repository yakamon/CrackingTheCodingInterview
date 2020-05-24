package circulation

import "testing"

func TestDetectCirculaton(t *testing.T) {
	valsHead := []interface{}{1, 2, 3, 4}
	valsCirculateStart := []interface{}{5}
	valsTail := []interface{}{6, 7, 8, 9, 10}

	// Create circulating linked list
	listHead := NewBySlice(valsHead)
	listCirculateStart := NewBySlice(valsCirculateStart)
	listTail := NewBySlice(valsTail)
	test1List := listHead
	test1List.SetNext(listCirculateStart)
	test1List.Tail().SetNext(listTail)
	test1List.Tail().SetNext(listCirculateStart)
	node1, isCirculating1 := DetectCirculation(test1List)
	if !isCirculating1 {
		t.Error("This list should be circular.")
	}
	if node1 != listCirculateStart {
		t.Errorf("First node of circulating list is incorrect.\nActual  : %v\nExpected: %v", node1, listCirculateStart)
	}

	// Create not circulating list
	test2List := NewBySlice(valsHead)
	node2, isCirculating2 := DetectCirculation(test2List)
	if isCirculating2 {
		t.Error("This list should not be circular.")
	}
	if node2 != nil {
		t.Error("The node should be nil")
	}
}

func BenchmarkDetectCirculation(b *testing.B) {
	valsHead := []interface{}{1, 2, 3, 4}
	valsCirculateStart := []interface{}{5}
	valsTail := []interface{}{6, 7, 8, 9, 10}
	listHead := NewBySlice(valsHead)
	listCirculateStart := NewBySlice(valsCirculateStart)
	listTail := NewBySlice(valsTail)
	test1List := listHead
	test1List.SetNext(listCirculateStart)
	test1List.Tail().SetNext(listTail)
	test1List.Tail().SetNext(listCirculateStart)

	for i := 0; i < b.N; i++ {
		b.Log(DetectCirculation(test1List))
	}
}

func TestDetectCirculaton2(t *testing.T) {
	valsHead := []interface{}{1, 2, 3, 4}
	valsCirculateStart := []interface{}{5}
	valsTail := []interface{}{6, 7, 8, 9, 10}

	// Create circulating linked list
	listHead := NewBySlice(valsHead)
	listCirculateStart := NewBySlice(valsCirculateStart)
	listTail := NewBySlice(valsTail)
	test1List := listHead
	test1List.SetNext(listCirculateStart)
	test1List.Tail().SetNext(listTail)
	test1List.Tail().SetNext(listCirculateStart)
	node1, isCirculating1 := DetectCirculation2(test1List)
	if !isCirculating1 {
		t.Error("This list should be circular.")
	}
	if node1 != listCirculateStart {
		t.Errorf("First node of circulating list is incorrect.\nActual  : %v\nExpected: %v", node1, listCirculateStart)
	}

	// Create not circulating list
	test2List := NewBySlice(valsHead)
	node2, isCirculating2 := DetectCirculation2(test2List)
	if isCirculating2 {
		t.Error("This list should not be circular.")
	}
	if node2 != nil {
		t.Error("The node should be nil")
	}
}

func BenchmarkDetectCirculation2(b *testing.B) {
	valsHead := []interface{}{1, 2, 3, 4}
	valsCirculateStart := []interface{}{5}
	valsTail := []interface{}{6, 7, 8, 9, 10}
	listHead := NewBySlice(valsHead)
	listCirculateStart := NewBySlice(valsCirculateStart)
	listTail := NewBySlice(valsTail)
	test1List := listHead
	test1List.SetNext(listCirculateStart)
	test1List.Tail().SetNext(listTail)
	test1List.Tail().SetNext(listCirculateStart)

	for i := 0; i < b.N; i++ {
		b.Log(DetectCirculation2(test1List))
	}
}
