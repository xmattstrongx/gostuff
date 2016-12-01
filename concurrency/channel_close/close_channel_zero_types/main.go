package main

import "fmt"

func main() {
	done := make(chan int)

	go func() {
		done <- 4
	}()

	fmt.Println(<-done)
	close(done)

	// can still write from closed channel but it will write default zero value for chan type
	fmt.Println(<-done)
}
