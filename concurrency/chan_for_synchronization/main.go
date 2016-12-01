package main

func main() {
	done := make(chan bool)

	go func() {
		println("goroutine message")

		// We are only intersted in the fact of sending itself,
		// but not in the data being sent
		done <- true
	}()

	println("main function message")
	<-done
}
