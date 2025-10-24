package arrays

import "sort"

// TwoSum ищет два числа, сумма которых равна target и возвращает их индексы
func TwoSum(nums []int, target int) []int {
	data := map[int]int{}

	for i, n := range nums {
		diff := target - n

		if _, ok := data[diff]; ok {
			return []int{data[diff], i}
		}

		data[n] = i
	}

	return []int{}
}

// TwoSumWithoutMap ищет два числа, сумма которых равна target и возвращает эти числа.
func TwoSumWithoutMap(nums []int, target int) []int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	l := 0
	r := len(nums) - 1

	for l != r {
		start := nums[l]
		end := nums[r]
		sum := start + end

		if sum == target {
			return []int{start, end}
		} else if sum > target {
			r--
		} else if sum < target {
			l++
		}
	}

	return []int{}
}
