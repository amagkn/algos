package arrays

/*
ArrayDiff

https://www.codewars.com/kata/523f5d21c841566fde000009/train/go
*/
func ArrayDiff(a, b []int) []int {
	data := make(map[int]struct{})

	for _, n := range b {
		data[n] = struct{}{}
	}

	var res []int
	for _, n := range a {
		if _, ok := data[n]; !ok {
			res = append(res, n)
		}
	}

	return res
}
