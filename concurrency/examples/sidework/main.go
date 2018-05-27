package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of program")
	quitMsg := make(chan string)

	go doSideWork(quitMsg)

	go func() {
		fmt.Println("Start of inline side work")
		time.Sleep(2 * time.Second)
		quitMsg <- "End of inline side work"
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("Executing main work: %d\n", i+1)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Finished main work: %d\n", i+1)
	}

	for {
		select {
		case msg := <-quitMsg:
			fmt.Println(msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Time over")
			return
		}
	}
}

func doSideWork(quitMsg chan<- string) {
	fmt.Println("Start of side work")
	time.Sleep(2 * time.Second)
	quitMsg <- "End of side work"
}
