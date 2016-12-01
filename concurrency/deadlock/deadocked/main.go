package main

func main() {
	c := make(chan int)

	//deadlocked because the code writes to and reads from the channel in a single thread.
	c <- 42
	val := <-c
	println(val)
}

// output
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
//      /fullpathtofile/channelsio.go:5 +0x54
// exit status 2
