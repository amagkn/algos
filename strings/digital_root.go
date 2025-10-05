package strings

import "strconv"

/*
DigitalRoot

https://www.codewars.com/kata/541c8630095125aba6000c00/train/go
*/
func DigitalRoot(n int) int {
	str := strconv.Itoa(n)

	for len(str) > 1 {
		var curr int

		for _, v := range str {
			x, _ := strconv.Atoi(string(v))
			curr += x
		}

		str = strconv.Itoa(curr)
	}

	res, _ := strconv.Atoi(str)

	return res
}
