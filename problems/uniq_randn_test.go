package problems

import "testing"

func isUnique(nums []int) bool {
	seen := map[int]struct{}{}

	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return false
		}

		seen[n] = struct{}{}
	}

	return true
}

func TestUniqRandn(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedLength int
	}{
		{
			name:           "нулевая длина",
			input:          0,
			expectedLength: 0,
		},
		{
			name:           "слайс определенной длины с уникальными числами",
			input:          5,
			expectedLength: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := UniqRandn(tt.input)

			if len(res) != tt.expectedLength {
				t.Errorf("UniqRandn(%d) = длина слайса %v, ожидалась длина слайса %d", tt.input, len(res), tt.expectedLength)
			}

			if !isUnique(res) {
				t.Errorf("UniqRandn(%d) = %v, ожидали слайс уникальных чисел", tt.input, res)
			}
		})
	}
}
