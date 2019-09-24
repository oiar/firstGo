package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	go func() {
		messages <- "1"
		messages <- "2"
	}()

	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
