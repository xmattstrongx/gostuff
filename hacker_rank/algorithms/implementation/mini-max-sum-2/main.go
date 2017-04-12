package main

import "fmt"

func main() {
	numberList := []int{0, 1, 4, 2, 3}

	var least, greatest, sum int

	for val, _ := range numberList {
		if val < least {
			least = val
		}
		if val > greatest {
			greatest = val
		}
		sum += val
	}
	fmt.Printf("greatest: %d\n", greatest)
	fmt.Printf("least: %d\n", least)

	fmt.Printf("min: %d\n", sum-greatest)
	fmt.Printf("max: %d\n", sum-least)
}
