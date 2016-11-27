package main

import (
	"fmt"
	"math"
)

func main() {
	size, shiftLength := readInSizeAndShiftLength()
	numbers := readInArray(size)
	shiftedNumbers := getArrayAfterShift(numbers, shiftLength)
	printArray(shiftedNumbers)
}

func readInSizeAndShiftLength() (size int, shiftLength int) {
	fmt.Scanf("%d%d", &size, &shiftLength)
	return size, shiftLength
}

func readInArray(size int) []int {
	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &numbers[i])
	}

	return numbers
}

func getArrayAfterShift(numbers []int, shiftLength int) []int {
	numbersLength := len(numbers)
	shiftedNumbers := make([]int, len(numbers))

	for i, v := range numbers {
		if i-shiftLength >= 0 {
			shiftedNumbers[i-shiftLength] = v
		} else {
			absoluteDifference := int(math.Abs(float64(i - shiftLength)))
			shiftedNumbers[numbersLength-absoluteDifference] = v
		}
	}
	return shiftedNumbers
}

func printArray(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d ", numbers[i])
	}
}
