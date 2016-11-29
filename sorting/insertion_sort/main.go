package main

import "fmt"

func main() {

	numberList := []int{6, 5, 1, 3, 8, 4, 7, 9, 2}
	fmt.Println(numberList)

	orderedNumberList := insertionSort(numberList)
	fmt.Println(orderedNumberList)

}

func insertionSort(numberList []int) []int {
	end := len(numberList)
	for i := 1; i < end; i++ {
		for j := i; j > 0 && numberList[j] < numberList[j-1]; j-- {
			numberList[j], numberList[j-1] = numberList[j-1], numberList[j]
			j = j - 1
		}
	}

	return numberList
}
