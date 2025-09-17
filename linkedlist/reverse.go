package linkedlist

func Reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func ClearReverse(head *ListNode) *ListNode {
	var newHead *ListNode
	curr := head

	for curr != nil {
		newNode := &ListNode{Val: curr.Val}
		newNode.Next = newHead
		newHead = newNode
		curr = curr.Next
	}

	return newHead
}
