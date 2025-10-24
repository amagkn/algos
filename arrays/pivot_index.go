package arrays

/*
PivotIndex

https://leetcode.com/problems/find-pivot-index/
*/
func PivotIndex(nums []int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}

	var currSum int
	for i, n := range nums {
		rightSum := sum - n - currSum

		if rightSum == currSum {
			return i
		}

		currSum += n
	}

	return -1
}
