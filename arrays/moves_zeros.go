package arrays

/*
MovesZeros

MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }

https://leetcode.com/problems/move-zeroes/
*/
func MovesZeros(arr []int) []int {
	pos := 0
	for i, v := range arr {
		if v != 0 {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}

	return arr
}
