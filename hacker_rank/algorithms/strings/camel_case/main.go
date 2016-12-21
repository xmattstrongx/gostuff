package main

import (
	"fmt"
	"regexp"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	var capRegEx = regexp.MustCompile(`[A-Z]`)
	fmt.Println(len(capRegEx.FindAllStringSubmatch(input, -1)) + 1)
}
