package main

import (
	"fmt"
	"strings"
)

func main() {
	height := getStairCaseHeight()
	printStaircase(height)
}

func getStairCaseHeight() int {
	var height int
	fmt.Scanf("%d", &height)
	return height
}

func printStaircase(height int) {
	emptySpace := height - 1
	for i := 1; i <= height; i++ {
		space := strings.Repeat(" ", emptySpace)
		steps := strings.Repeat("#", i)
		fmt.Printf("%s%s\n", space, steps)
		emptySpace--
	}
}
