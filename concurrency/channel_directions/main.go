package main

import (
	"fmt"
	"time"
)

// The <- operator is on the right of the chan keyword
// This means this channel can only be sent to
// Attempting to receive from it will cause compilation error
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// The <- operator is on the right of the chan keyword
// This means this channel can only be sent to
// Attempting to receive from it will cause compilation error
func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// The <- operator is on the left of the chan keyword
// This means this channel can only receive from a channel
// Attempting to send to it will cause compilation error
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
