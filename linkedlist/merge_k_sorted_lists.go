package linkedlist

func MergeKSortedLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{}

	// Пока список не будет пуст - итерируемся
	curr := dummy
	for len(lists) != 0 {
		// Ищем ноду, в которой наименьшее текущее значение
		nodeWithMinValueIndex := 0
		for i, l := range lists {
			if l.Val < lists[nodeWithMinValueIndex].Val {
				nodeWithMinValueIndex = i
			}
		}

		// Добавляем ноду с наименьшим значением в результат
		curr.Next = lists[nodeWithMinValueIndex]
		curr = curr.Next

		// Если у этого списка больше не осталось нод - удаляем его из массива списка, иначе - просто инкриментим его
		if lists[nodeWithMinValueIndex].Next == nil {
			lists = append(lists[:nodeWithMinValueIndex], lists[nodeWithMinValueIndex+1:]...)
		} else {
			lists[nodeWithMinValueIndex] = lists[nodeWithMinValueIndex].Next
		}
	}

	return dummy.Next
}
