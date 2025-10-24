package strings

func Reverse(str string) string {
	res := make([]rune, 0, len(str))
	runes := []rune(str)

	for i := len(runes) - 1; i >= 0; i-- {
		res = append(res, runes[i])
	}

	return string(res)
}
