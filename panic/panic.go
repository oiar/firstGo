package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in main from error:", err)
		}
		fmt.Println("Execution finished")
	}()

	go func() {
		panic("panic in goroutine")
	}()

	time.Sleep(1 * time.Second)
}
