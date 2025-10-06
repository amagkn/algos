package strings

import (
	"strings"
)

/*
SpinWords

https://www.codewars.com/kata/5264d2b162488dc400000001/train/go
*/
func SpinWords(str string) string {
	words := strings.Split(str, " ")

	for i, w := range words {
		runes := []rune(w)

		if len(runes) >= 5 {
			var builder strings.Builder

			for k := len(runes) - 1; k >= 0; k-- {
				builder.WriteRune(runes[k])
			}

			words[i] = builder.String()
		}
	}

	return strings.Join(words, " ")
}
