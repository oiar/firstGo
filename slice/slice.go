package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recover from error: %v", err)
		}
	}()
	s := make([]int, 4, 5)
	fmt.Print("len = %d, cap = %d", len(s), cap(s))
}
