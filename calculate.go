package main

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func Avg(tab []int) int {
	var sum int = 0

	for _, num := range tab {
		sum += num
	}

	return sum / len(tab)
}
