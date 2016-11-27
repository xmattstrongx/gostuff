package main

import (
	"fmt"
	"log"
)

func main() {
	alice, bob, err := readData()
	if err != nil {
		log.Fatal(err)
	}

	aliceResult, bobResult := calculateResults(alice, bob)
	fmt.Printf("%d %d", aliceResult, bobResult)
}

func readData() ([]int, []int, error) {

	alice := make([]int, 3)

	for i := 0; i < 3; i++ {
		_, err := fmt.Scanf("%d", &alice[i])
		if err != nil {
			return nil, nil, err
		}

	}

	bob := make([]int, 3)

	for i := 0; i < 3; i++ {
		_, err := fmt.Scanf("%d", &bob[i])
		if err != nil {
			return nil, nil, err
		}

	}

	return alice, bob, nil
}

func calculateResults(alice []int, bob []int) (int, int) {
	var aliceResult int
	var bobResult int
	for i := 0; i < len(alice); i++ {
		if alice[i] > bob[i] {
			aliceResult++
		} else if bob[i] > alice[i] {
			bobResult++
		}
	}
	return aliceResult, bobResult
}
