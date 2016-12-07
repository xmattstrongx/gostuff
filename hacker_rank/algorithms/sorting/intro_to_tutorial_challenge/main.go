package main

import "fmt"

type numberList struct {
	Value      int
	Size       int
	NumberList []int
}

func main() {

	n := numberList{}
	n.readInData()
	found := n.findValueLocation()
	fmt.Println(found)
}

func (n *numberList) readInData() {
	fmt.Scanf("%d\n%d", &n.Value, &n.Size)

	n.NumberList = make([]int, n.Size)
	for i := 0; i < n.Size; i++ {
		fmt.Scanf("%d", &n.NumberList[i])
	}
}

func (n *numberList) findValueLocation() int {
	for i := 0; i < len(n.NumberList); i++ {
		if n.NumberList[i] == n.Value {
			return i
		}
	}
	return 0
}
