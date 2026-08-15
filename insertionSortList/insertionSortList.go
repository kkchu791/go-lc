package insertionsortlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortList(head *ListNode) *ListNode {
	sentinel := &ListNode{}

	for head != nil {
		tmp := sentinel

		for tmp.Next != nil && tmp.Next.Val < head.Val {
			tmp = tmp.Next
		}

		// insert copy of head into sorted position

		curr := &ListNode{
			Val:  head.Val,
			Next: tmp.Next,
		}

		tmp.Next = curr

		head = head.Next
	}

	return sentinel.Next

}
