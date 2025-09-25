package linkedlist

func Create(values []int) *ListNode {
	dummy := &ListNode{}

	curr := dummy

	for _, v := range values {
		curr.Next = &ListNode{
			Val: v,
		}
		curr = curr.Next
	}

	return dummy.Next
}
