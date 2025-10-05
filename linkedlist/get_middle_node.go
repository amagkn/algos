package linkedlist

/*
GetMiddleNode

https://leetcode.com/problems/middle-of-the-linked-list/
*/
func GetMiddleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
