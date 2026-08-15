package insertionsortlist

func insertionSortList(head *ListNode) *ListNode {
	sortedRes := &ListNode{}
	curr := head

	for curr != nil {
		prev := sortedRes

		for prev.Next != nil && prev.Next.Val <= curr.Val {
			prev = prev.Next
		}

		next := curr.Next

		// insert the current node to the new list
		curr.Next = prev.Next

		prev.Next = curr

		curr = next
	}

	return sortedRes.Next
}
