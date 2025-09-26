package main

import (
	"fmt"

	"github.com/amagkn/algos/linkedlist"
)

func main() {
	head := linkedlist.Create([]int{1, 2, 3, 4})

	fmt.Println(linkedlist.ToSlice(linkedlist.Reorder(head)))
}
