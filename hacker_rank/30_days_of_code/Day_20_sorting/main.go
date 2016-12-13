package main

import "fmt"

// type BubbleSort interface {
// 	bubbleSort() []int
// }

type BubbleValues struct {
	size       int
	numberList []int
}

type SortedNumberList struct {
	swaps            int
	sortedNumberList []int
}

func main() {
	bubbleValues := BubbleValues{}
	bubbleValues.readInData()

}

func (b *BubbleValues) readInData() {
	b.getSize()
	b.getNumberList()

	sorted := b.bubbleSort()
	sorted.printOutcome()
}

func (b *BubbleValues) getSize() {
	fmt.Scanf("%d", &b.size)
}

func (b *BubbleValues) getNumberList() {
	b.numberList = make([]int, b.size)
	for i := 0; i < b.size; i++ {
		fmt.Scanf("%d", &b.numberList[i])
	}
}

func (b *BubbleValues) bubbleSort() SortedNumberList {
	snl := SortedNumberList{}
	snl.sortedNumberList = b.numberList
	var temp int
	for i := 0; i < b.size; i++ {

		for j := 0; j < b.size-1; j++ {
			if snl.sortedNumberList[j] > snl.sortedNumberList[j+1] {
				temp = snl.sortedNumberList[j+1]
				snl.sortedNumberList[j+1] = snl.sortedNumberList[j]
				snl.sortedNumberList[j] = temp
				snl.swaps++
			}
		}
	}
	return snl
}

func (s *SortedNumberList) printOutcome() {
	fmt.Printf("Array is sorted in %d swaps.\n", s.swaps)
	fmt.Printf("First Element: %d \n", s.sortedNumberList[0])
	fmt.Printf("Last Element: %d", s.sortedNumberList[len(s.sortedNumberList)-1])
}
