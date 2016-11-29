package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d", &size)

	numberList := readInputToArray(size)

	printReversedArray(numberList)
}

func readInputToArray(size int) []int {
	numberList := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &numberList[i])
	}

	return numberList
}

func printReversedArray(numberList []int) {
	for i := len(numberList) - 1; i >= 0; i-- {
		fmt.Printf("%d ", numberList[i])
	}
}
