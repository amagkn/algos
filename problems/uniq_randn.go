package problems

import "math/rand"

// UniqRandn генерирует слайс n уникальных чисел
func UniqRandn(n int) []int {
	if n <= 0 {
		return []int{}
	}

	result := make([]int, n)
	seen := make(map[int]struct{})

	for i := 0; i < n; {
		num := rand.Intn(10)

		if _, ok := seen[num]; !ok {
			result[i] = num
			seen[num] = struct{}{}
			i++
		}
	}

	return result
}
