package strings

import "fmt"

type SymCounter struct {
	sym   rune
	count int
}

func RLE(s string) string {
	result := ""

	var curr *SymCounter
	for _, sym := range s {
		if curr == nil {
			curr = &SymCounter{
				sym:   sym,
				count: 1,
			}
		} else if curr.sym == sym {
			curr.count++
		} else if curr.sym != sym {
			if curr.count > 1 {
				result += fmt.Sprintf("%s%d", string(curr.sym), curr.count)
			} else {
				result += fmt.Sprintf("%s", string(curr.sym))
			}

			curr = &SymCounter{
				sym:   sym,
				count: 1,
			}
		}
	}

	if curr.count > 1 {
		result += fmt.Sprintf("%s%d", string(curr.sym), curr.count)
	} else {
		result += fmt.Sprintf("%s", string(curr.sym))
	}

	return result
}
