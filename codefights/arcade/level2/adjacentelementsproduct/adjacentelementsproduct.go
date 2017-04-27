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

// Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.

// Example

// For inputArray = [3, 6, -2, -5, 7, 3], the output should be
// adjacentElementsProduct(inputArray) = 21.

// 7 and 3 produce the largest product.

// Input/Output

// [time limit] 4000ms (go)
// [input] array.integer inputArray

// An array of integers containing at least two elements.

// Guaranteed constraints:
// 2 ≤ inputArray.length ≤ 10,
// -1000 ≤ inputArray[i] ≤ 1000.

// [output] integer

// The largest product of adjacent elements.
