package strings

import "strings"

// CountWords считает количество слов в строке.
func CountWords(str string) map[string]int {
	counter := map[string]int{}

	words := strings.Split(str, " ")

	for _, w := range words {
		if len(w) == 0 {
			continue
		}

		counter[w]++
	}

	return counter
}

func CountWordsWithoutSplit(str string) map[string]int {
	res := map[string]int{}

	currentStr := ""
	for _, r := range str {
		if string(r) == " " {
			if len(currentStr) > 0 {
				res[currentStr]++
				currentStr = ""
			}
			continue
		}

		currentStr += string(r)
	}

	if len(currentStr) > 0 {
		res[currentStr]++
	}

	return res
}
