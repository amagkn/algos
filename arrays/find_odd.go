package arrays

/*
FindOdd

https://www.codewars.com/kata/54da5a58ea159efa38000836/train/go
*/
func FindOdd(seq []int) int {
	data := make(map[int]struct{})

	for _, n := range seq {
		if _, ok := data[n]; ok {
			delete(data, n)
		} else {
			data[n] = struct{}{}
		}
	}

	for n := range data {
		return n
	}

	return 0
}
