package linkedlist

/*
PalindromeLinkedList

Можно  полностью перевернуть список и сравнить перевернутый список с оригиналом, но это не совсем оптимально.
Лучше сравнить перевернутую вторую "половину" листа с первой половиной (не перевернутой), для этого придется искать середину списка.
*/

func PalindromeLinkedList(head *ListNode) bool {
	// Ищем середину листа, это будет первый узел во второй половине листа
	secondHalfBeginNode := GetMiddleNode(head)
	// Переворачиваем вторую часть листа
	reversedSecondHalfBeginNode := Reverse(secondHalfBeginNode)

	// Сравниваем обе половины листа
	p1 := head
	p2 := reversedSecondHalfBeginNode
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}

		p2 = p2.Next
		p1 = p1.Next
	}

	return true
}

func PalindromeLinkedListMySolution(head *ListNode) bool {
	// Ищем середину
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// (Не обязательно!) Если fast не указывает на nil, значит элементов нечетное количество, и slow указывает на центральный,
	// не включаем его в вторую часть и тем самым он вообще не будет учавствать в сравнении элементов
	//if fast != nil {
	//	slow = slow.Next
	//}

	// Переворачиваем вторую часть (делаем reverse)
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Итерируемся по обеим частям, но цикл завершится когда закончится именно вторая часть, потому-что первая часть включает и вторую
	reversedSecondPart := prev
	firstPart := head
	for reversedSecondPart != nil {
		if reversedSecondPart.Val != firstPart.Val {
			return false
		}

		reversedSecondPart = reversedSecondPart.Next
		firstPart = firstPart.Next
	}

	return true
}
