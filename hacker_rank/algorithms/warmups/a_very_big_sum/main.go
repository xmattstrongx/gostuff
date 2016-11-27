package main

import (
	"fmt"
	"log"
)

func main() {
	var length int
	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		log.Fatal(err)
	}

	var currentAmount, total int
	for i := 0; i < length; i++ {
		_, err := fmt.Scanf("%d", &currentAmount)
		if err != nil {
			log.Fatal(err)
		}
		total += currentAmount
	}

	fmt.Println(total)
}
