package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(Add(1, 2))
	fmt.Println(Sub(5, 2))
	fmt.Println(Mul(2, 3))
	fmt.Println(Div(6, 2))
	fmt.Println(Avg(numbers))
}
