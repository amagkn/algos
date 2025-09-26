package linkedlist

func Reorder(head *ListNode) *ListNode {
	// Ищем середину, стандартный подход
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Берем следующий после середины
	second := slow.Next
	// и явно разделяем список на две части
	slow.Next = nil

	// Реверс второй части списка
	var prev *ListNode
	curr := second
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Пробегаемся по первой и по перевернутой второй части и мержим через один
	secondReversed := prev
	first := head
	for secondReversed != nil {
		tmp1 := first.Next
		tmp2 := secondReversed.Next

		first.Next = secondReversed
		secondReversed.Next = tmp1

		first = tmp1
		secondReversed = tmp2
	}

	return head
}
