package main

import "fmt"

func main() {

	numberList := []int{6, 5, 1, 3, 8, 4, 7, 9, 2}
	fmt.Println(numberList)

	orderedNumberList := bubbleSort(numberList)
	fmt.Println(orderedNumberList)

}

func bubbleSort(numberList []int) []int {
	end := len(numberList)
	for i := end; i >= 1; i-- {
		for j := 0; j < i-1; j++ {
			if numberList[j] > numberList[j+1] {
				numberList[j], numberList[j+1] = numberList[j+1], numberList[j]
			}
		}
	}

	return numberList
}
