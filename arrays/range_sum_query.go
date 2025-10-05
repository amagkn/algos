package arrays

type NumArray struct {
	prefix []int
}

func NewNumArray(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	return NumArray{prefix: prefix}
}

/*
RangeSumQuery

# Нужно решать с помощью префиксного массива

https://leetcode.com/problems/range-sum-query-immutable/
*/
func (n *NumArray) RangeSumQuery(left int, right int) int {
	return n.prefix[right+1] - n.prefix[left]
}
