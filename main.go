package main

import (
	"fmt"

	"github.com/amagkn/algos/linkedlist"
)

func main() {
	lists := []*linkedlist.ListNode{
		linkedlist.Create([]int{1, 5, 10}),
		linkedlist.Create([]int{2, 7, 12}),
		linkedlist.Create([]int{4, 6, 22}),
	}

	fmt.Println(linkedlist.ToSlice(linkedlist.MergeKSortedLists(lists))) // [1,2,4,5,6,7,10,12,22]
}
