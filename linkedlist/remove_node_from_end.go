package linkedlist

func RemoveNodeFromEnd(head *ListNode, idx int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	curr := dummy
	var count int
	for curr != nil {
		count++
		curr = curr.Next
	}

	target := count - idx - 1
	curr = dummy
	for range target {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next

	return dummy.Next
}
