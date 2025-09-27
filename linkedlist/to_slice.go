package linkedlist

func ToSlice(head *ListNode) []any {
	var res []any

	curr := head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	return res
}
