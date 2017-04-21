package anomalous

import (
	"math"
	"fmt"
	"strconv"
)

func Gongyue(m, n int) (num int) {
	if m % n == 0 {
		num = n
	} else {
		m, n = n, m % n
		num = Gongyue(m, n)
	}
	return
}

func Hanoi(n int, x, y, z string) {
	if n == 1 {
		fmt.Println(x + "->" + z)
	} else {
		Hanoi(n - 1, x, z, y)
		fmt.Println(x + "->" + z)
		Hanoi(n - 1, y, x, z)
	}
}

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values) - 1
	for head < tail {
		fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head + 1:])
}

func Fibonacci_sequence(n int) (sum int) {
	if n == 1 || n == 2 {
		sum = 1
	} else {
		sum = Fibonacci_sequence(n - 1) + Fibonacci_sequence(n - 2)
	}
	return
}

func Prime(min, max int) {
	for min <= max {
		status := 0
		for i := 1; i <= min; i++ {
			if min % i == 0 {
				status++
			}
			if status == 3 {
				break
			}
		}
		if status == 3 {
			min++
			continue
		}
		fmt.Println(min)
		min++
	}
}

func Narcissus(min, max int) {
	for min <= max {
		b := float64(min / 100)
		by := min % 100
		s := float64(by / 10)
		g := float64(min % 10)
		if math.Pow(b, 3) + math.Pow(s, 3) + math.Pow(g, 3) == float64(min) {
			fmt.Println(min)
		}
		min++
	}
}

func Resolved(num int) (str string) {
	for i := 2; i < num; i++ {
		if num % i == 0 {
			str += strconv.Itoa(i) + "*" + Resolved(num / i)
			break
		}
	}
	if len(str) == 0 {
		str += strconv.Itoa(num)
	}
	return
}

func Qiu(h, n, s float64) (height, sum float64) {
	if n == 11 {
		height, sum = h, s
	} else {
		height, sum = Qiu(0.5 * h, n + 1, s + h)
	}
	return
}
