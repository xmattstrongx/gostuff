package main

import (
	"fmt"
	"log"
)

func main() {
	data, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumData(data))

}

func readData() ([]int, error) {
	var length int

	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		return nil, err
	}

	data := make([]int, length)

	for i := range data {
		_, err := fmt.Scanf("%d", &data[i])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func sumData(data []int) int {
	var total int
	for _, v := range data {
		total += v
	}
	return total
}
