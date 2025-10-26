package strings

import "testing"

func TestRLE(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Базовая проверка",
			input:    "AAAABBBCCXYZDDDDEEEFFFFAAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBBBB",
			expected: "A4B3C2XYZD4E3F4A7B30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RLE(tt.input)

			if result != tt.expected {
				t.Errorf("RLE(%s) = %s, expected %s", tt.input, result, tt.expected)
			}
		})
	}
}
