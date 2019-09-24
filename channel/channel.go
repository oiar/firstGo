package main

import "fmt"

func main() {
	messages := make(chan string, 1)

	messages <- "ping"

	msg := <-messages
	fmt.Println(msg)
}
