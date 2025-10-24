package strings

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:     "простая строка с разными словами",
			input:    "hello world",
			expected: map[string]int{"hello": 1, "world": 1},
		},
		{
			name:     "повторяющиеся слова",
			input:    "hello hello world",
			expected: map[string]int{"hello": 2, "world": 1},
		},
		{
			name:     "одно слово",
			input:    "hello",
			expected: map[string]int{"hello": 1},
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "строка с множественными пробелами",
			input:    "hello  world",
			expected: map[string]int{"hello": 1, "world": 1},
		},
		{
			name:     "строка только из пробелов",
			input:    "   ",
			expected: map[string]int{},
		},
		{
			name:     "строка с одним пробелом",
			input:    " ",
			expected: map[string]int{},
		},
		{
			name:     "слова с разным регистром",
			input:    "Hello hello HELLO",
			expected: map[string]int{"Hello": 1, "hello": 1, "HELLO": 1},
		},
		{
			name:     "много повторяющихся слов",
			input:    "test test test test",
			expected: map[string]int{"test": 4},
		},
		{
			name:     "слова с знаками препинания",
			input:    "hello, world!",
			expected: map[string]int{"hello,": 1, "world!": 1},
		},
		{
			name:     "русские слова",
			input:    "привет мир привет",
			expected: map[string]int{"привет": 2, "мир": 1},
		},
		{
			name:     "числа как слова",
			input:    "123 456 123",
			expected: map[string]int{"123": 2, "456": 1},
		},
		{
			name:     "пробелы в начале и конце",
			input:    " hello world ",
			expected: map[string]int{"hello": 1, "world": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountWordsWithoutSplit(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountWords(%s) = %v, ожидалось %v", tt.input, result, tt.expected)
			}
		})
	}
}
