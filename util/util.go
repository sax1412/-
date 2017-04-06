package util

func Substr(str string, start int, length int) string {
	//rs := []rune(str)
	rl := len(str)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(str[start:end])
}

func Str_delete(s string) string {
	status := 0
	var str []rune
	for _, r := range s {
		if r == rune('<') {
			status = 1
			continue
		}
		if r == rune('>') {
			status = 0
			continue
		}
		if status == 0 {
			str = append(str, r)
		}
	}
	return string(str)
}
