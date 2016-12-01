package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d", &input)

	fmt.Println(factorial(input))
}

func factorial(input int) int {
	if input <= 1 {
		return 1
	}
	return input * factorial(input-1)
}
