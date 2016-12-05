package main

import "fmt"

func main() {
	size := getSize()
	numberList := readInData(size)

	positives, negatives, zeros := getAverages(numberList, size)
	fmt.Printf("%f\n%f\n%f", positives, negatives, zeros)
}

func getSize() int {
	var length int
	fmt.Scanf("%d", &length)
	return length
}

func readInData(size int) []int {
	numberList := make([]int, size)
	for i := range numberList {
		fmt.Scanf("%d", &numberList[i])
	}
	return numberList
}

func getAverages(numberList []int, size int) (float64, float64, float64) {
	positives, negatives, zeros := getAmountsOfPositivesZeroesAndNegatives(numberList)
	return positives / float64(size), negatives / float64(size), zeros / float64(size)
}

func getAmountsOfPositivesZeroesAndNegatives(numberList []int) (float64, float64, float64) {
	var positives, negatives, zeros float64
	for _, num := range numberList {
		if num == 0 {
			zeros++
		} else if num > 0 {
			positives++
		} else if num < 0 {
			negatives++
		}
	}
	return positives, negatives, zeros
}
