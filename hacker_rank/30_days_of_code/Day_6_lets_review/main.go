package main

import "fmt"

func main() {
	var count int

	fmt.Scanf("%d", &count)

	inputList := readInStrings(count)

	printCharacters(inputList)
}

func readInStrings(count int) []string {
	inputList := make([]string, count)
	for i := 0; i < count; i++ {
		fmt.Scanf("%s", &inputList[i])
	}

	return inputList
}

func printCharacters(list []string) {
	for _, v := range list {
		printEvens(v)
		fmt.Printf(" ")
		printOdds(v)
		fmt.Printf("\n")
	}
}

func printEvens(word string) {
	for i, v := range word {
		if i%2 == 0 {
			c := string(v)
			fmt.Printf("%s", c)
		}
	}
}

func printOdds(word string) {
	for i, v := range word {
		if i%2 == 1 {
			c := string(v)
			fmt.Printf("%s", c)
		}
	}
}
