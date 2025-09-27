package main

import (
	"fmt"

	"github.com/amagkn/algos/linkedlist"
)

func main() {
	l1 := linkedlist.Create([]int{1, 3, 3, 5})
	l2 := linkedlist.Create([]int{2, 4, 6, 6, 7, 8, 10})

	fmt.Println(linkedlist.ToSlice(linkedlist.MergeTwoSortedLists(l1, l2))) // [1 2 3 3 4 5 6 6 7 8 10]
}
