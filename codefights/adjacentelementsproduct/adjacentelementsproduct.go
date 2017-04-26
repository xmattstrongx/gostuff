package adjacentelementsproduct

func adjacentElementsProduct(inputArray []int) int {
	var largest = -1000
	for i := 0; i < len(inputArray)-1; i++ {
		if val := inputArray[i] * inputArray[i+1]; val > largest {
			largest = val
		}
	}
	return largest
}
