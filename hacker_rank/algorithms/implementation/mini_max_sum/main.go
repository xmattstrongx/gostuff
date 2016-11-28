package main

import (
	"fmt"
	"sort"
)

func main() {
	var numberList int64arr
	numberList = readNumberList()

	sort.Sort(int64arr(numberList))

	least := calculateLeast(numberList)
	greatest := calculateGreatest(numberList)

	fmt.Printf("%d %d", least, greatest)
}

func readNumberList() []int64 {
	numberList := make([]int64, 5)

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &numberList[i])
	}

	return numberList
}

func calculateLeast(numberList []int64) int64 {
	var total int64
	for i := 0; i < 4; i++ {
		total += numberList[i]
	}
	return total
}

func calculateGreatest(numberList []int64) int64 {
	var total int64
	for i := 1; i < 5; i++ {
		total += numberList[i]
	}
	return total
}

// Implement sort by length interface for int64
type int64arr []int64

func (a int64arr) Len() int {
	return len(a)
}

func (a int64arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a int64arr) Less(i, j int) bool {
	return a[i] < a[j]
}
