package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)

	// Declare second integer, double, and String variables.
	var j uint32
	var f float32
	var t string

	// Read and save an integer, double, and String to your variables.
	fmt.Scan(&j, &f)
	b, _ := scanner.ReadString('\n')
	t = string(b)

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + j)

	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", (d + f))

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + t)
}
