package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string) // no buffer
	// message := make(chan string, 2) // buffered channel
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("send message start")
			message <- fmt.Sprintf("message %d", i)
			fmt.Println("send message end")
		}
	}()

	time.Sleep(time.Second * 3)

	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
}
