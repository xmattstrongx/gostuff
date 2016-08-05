package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%5 == 0 {
			if i%3 == 0 {
				fmt.Println("fizzbuzz")
			} else {
				fmt.Println("buzz")
			}
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}
	}
}
