package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numberList := []int{6, 5, 1, 3, 8, 4, 7, 9, 2}

	a := quicksort(numberList)
	fmt.Println(a)

}

func quicksort(numberList []int) []int {
	if len(numberList) < 2 {
		return numberList
	}
	wall, end := 0, len(numberList)-1
	pivotIndex := rand.Int() % len(numberList)
	numberList[pivotIndex], numberList[end] = numberList[end], numberList[pivotIndex]

	for i := range numberList {
		if numberList[end] > numberList[i] {
			numberList[i], numberList[wall] = numberList[wall], numberList[i]
			wall++
		}
	}

	numberList[wall], numberList[end] = numberList[end], numberList[wall]

	quicksort(numberList[:wall])
	quicksort(numberList[wall+1:])

	return numberList
}
