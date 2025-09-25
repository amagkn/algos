package main

import (
	"fmt"

	"github.com/amagkn/algos/linkedlist"
)

func main() {
	head := linkedlist.Create([]int{1, 2, 3, 3, 2, 1})

	fmt.Println(linkedlist.PalindromeLinkedList(head))
}
