package main

func main() {
	done := make(chan bool)
	go func() {
		println("goroutine message")

		// Just send a signal "I'm done"
		close(done)
	}()

	println("main function message")
	<-done
}
