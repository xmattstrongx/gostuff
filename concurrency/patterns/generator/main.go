package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator is a function that returns a channel

func main() {
	c := boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) //Receive expression is just a value
	}
	fmt.Println("You're boring; I'm leaving.")
}

func boring(msg string) <-chan string { // returns receive-only channel of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller
}
