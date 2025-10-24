package arrays

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "простой случай с положительными числами",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1}, // nums[0] + nums[1] = 2 + 7 = 9
		},
		{
			name:     "числа в середине массива",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2}, // nums[1] + nums[2] = 2 + 4 = 6
		},
		{
			name:     "два одинаковых числа",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1}, // nums[0] + nums[1] = 3 + 3 = 6
		},
		{
			name:     "отрицательные числа",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4}, // nums[2] + nums[4] = -3 + (-5) = -8
		},
		{
			name:     "смесь положительных и отрицательных",
			nums:     []int{-3, 4, 3, 90},
			target:   0,
			expected: []int{0, 2}, // nums[0] + nums[2] = -3 + 3 = 0
		},
		{
			name:     "большие числа",
			nums:     []int{1000, 2000, 3000},
			target:   5000,
			expected: []int{1, 2}, // nums[1] + nums[2] = 2000 + 3000 = 5000
		},
		{
			name:     "нет решения",
			nums:     []int{1, 2, 3},
			target:   7,
			expected: []int{},
		},
		{
			name:     "пустой массив",
			nums:     []int{},
			target:   0,
			expected: []int{},
		},
		{
			name:     "один элемент",
			nums:     []int{5},
			target:   5,
			expected: []int{},
		},
		{
			name:     "ноль в массиве",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3}, // nums[0] + nums[3] = 0 + 0 = 0
		},
		{
			name:     "дубликаты в массиве",
			nums:     []int{1, 1, 1, 1},
			target:   2,
			expected: []int{0, 1}, // любая пара индексов подойдет
		},
		{
			name:     "решение в начале массива",
			nums:     []int{5, 5, 10, 20},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "решение в конце массива",
			nums:     []int{1, 2, 3, 4, 5},
			target:   9,
			expected: []int{3, 4}, // nums[3] + nums[4] = 4 + 5 = 9
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := TwoSum(tt.nums, tt.target)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("TwoSum(%v, %d) = %v, ожидалось %v", tt.nums, tt.target, res, tt.expected)
			}
		})
	}
}

func TestTwoSumWithoutMap(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "простой случай с положительными числами",
			nums:     []int{2, 7, 11, 15},
			target:   18,
			expected: []int{7, 11},
		},
		{
			name:     "числа требуют сортировки",
			nums:     []int{3, 2, 4},
			target:   7,
			expected: []int{3, 4},
		},
		{
			name:     "два одинаковых числа",
			nums:     []int{3, 10, 12, 3, 1},
			target:   6,
			expected: []int{3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := TwoSumWithoutMap(tt.nums, tt.target)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("TwoSum(%v, %d) = %v, ожидалось %v", tt.nums, tt.target, res, tt.expected)
			}
		})
	}
}
