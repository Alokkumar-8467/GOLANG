package main

// When we have multiple goroutines running concurrently, we need a way to communicate between them. Channels provide a way for one goroutine to send data to another goroutine. Channels are a powerful feature of Go that allow us to build concurrent programs that are easy to reason about.


func main() {

	// How to create a channel in Go?
	messageChan := make(chan string)

		// How to send a message to a channel in Go?
	messageChan <- "Ping, channel!"

	// How to receive a message from a channel in Go?
	msg := <-messageChan
	println(msg)

}
