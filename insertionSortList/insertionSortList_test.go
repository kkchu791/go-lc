package insertionsortlist

import (
	"fmt"
	"testing"
)

func buildList(vals []int) *ListNode {
	sentinel := &ListNode{}
	tmp := sentinel
	for _, v := range vals {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return sentinel.Next
}

func listsEqual(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func TestInsertionSortList(t *testing.T) {
	ln := buildList([]int{4, 3, 5, 6})
	result := InsertionSortList(ln)
	expected := buildList([]int{3, 4, 5, 6})

	fmt.Println(result, expected)

	if listsEqual(result, expected) {
		fmt.Println("test passed")
	} else {
		t.Errorf("test did not pass")
	}
}
