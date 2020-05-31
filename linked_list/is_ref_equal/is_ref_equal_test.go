package refequal

import (
	"fmt"
	"testing"
)

func TestIsLinkedListEqual(t *testing.T) {
	// 連結リストとそのコピーを比較してみる
	testSlice := []interface{}{1, 2, 3, 4}
	eq1 := NewBySlice(testSlice)
	eq2 := eq1
	fmt.Println(eq1, eq2)
	fmt.Println(eq1 == eq2)
	fmt.Println(eq1.Next() == eq2.Next())

	// 同じ配列から作った、つまり同じ値を持つ、二つの連結リストを比較してみる
	neq1 := NewBySlice(testSlice)
	neq2 := NewBySlice(testSlice)
	fmt.Println(neq1, neq2)
	fmt.Println(neq1 == neq2)
	fmt.Println(neq1.Next() == neq2.Next())
}

func TestFindIntersecton(t *testing.T) {
	slice1 := []interface{}{1, 2, 3}
	slice2 := []interface{}{4, 5}
	sliceCommon := []interface{}{6, 7, 8, 9, 10}

	list1 := NewBySlice(slice1)
	list2 := NewBySlice(slice2)
	listCommon := NewBySlice(sliceCommon)

	// 共通部分があるパターン
	test1List1 := list1
	test1List1.Tail().SetNext(listCommon)
	test1List2 := list2
	test1List2.Tail().SetNext(listCommon)
	expected1 := listCommon
	if actual1, found := FindIntersection(test1List1, test1List2); !found || actual1 != expected1 {
		t.Errorf("FindIntersection does not work correctly.\nActual:   %v\nExpected: %v", actual1, expected1)
	}

	// 共通部分がないパターン
	test2List1 := NewBySlice(slice1)
	test2List2 := NewBySlice(slice2)
	if actual2, found := FindIntersection(test2List1, test2List2); found || actual2 != nil {
		t.Errorf("The return value should be {nil, false}\n")
	}
}

func BenchmarkFindIntersection(b *testing.B) {
	slice1 := []interface{}{1, 2, 3}
	slice2 := []interface{}{4, 5}
	sliceCommon := []interface{}{6, 7, 8, 9, 10}

	list1 := NewBySlice(slice1)
	list2 := NewBySlice(slice2)
	listCommon := NewBySlice(sliceCommon)

	test1List1 := list1
	test1List1.Tail().SetNext(listCommon)
	test1List2 := list2
	test1List2.Tail().SetNext(listCommon)

	for i := 0; i < b.N; i++ {
		b.Log(FindIntersection(test1List1, test1List2))
	}
}
