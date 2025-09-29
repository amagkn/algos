package linkedlist

func ToSlice(head *ListNode) []int {
	var res []int

	curr := head
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}

	return res
}
