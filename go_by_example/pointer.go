package main

import "fmt"

func normal(n int) {
	n = 0
}

func pointer(n *int) {
	*n = 0
}

func main() {
	n := 1
	normal(n)
	fmt.Println(n, "normal")
	pointer(&n)
	fmt.Println(n, "pointer")
	fmt.Println(&n, "pointer_address")
}
