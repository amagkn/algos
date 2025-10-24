package concurrency

import (
	"testing"
)

func TestSumSquares(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "числа как числа",
			input:    []int{1, 2, 3, 4},
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SumSquares(tt.input)

			if res != tt.expected {
				t.Errorf("SumSquares(%v) = %d, ожидалось %d", tt.input, res, tt.expected)
			}
		})
	}
}
