package linkedlist

func Create(length int) *ListNode {
	dummy := &ListNode{}

	curr := dummy
	for i := 0; i < length; i++ {
		curr.Next = &ListNode{
			Val: i + 1,
		}
		curr = curr.Next
	}

	return dummy.Next
}
