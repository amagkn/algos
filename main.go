package main

import (
	"fmt"

	"github.com/amagkn/algos/arrays"
)

func main() {
	numArray := arrays.NewNumArray([]int{-2, 0, 3, -5, 2, -1})

	fmt.Println(numArray.RangeSumQuery(0, 2)) // 1
	fmt.Println(numArray.RangeSumQuery(2, 5)) // -1
	fmt.Println(numArray.RangeSumQuery(0, 5)) // -3
}
