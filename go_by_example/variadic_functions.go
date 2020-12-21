package main

import "fmt"

func f(nums ...int) int {
	fmt.Print(nums)
	total := 0
	for a := range nums {
		fmt.Println(a, "a")
		total += a
	}
	fmt.Println(total," total")
	return total
}

func main() {
	nums := []int{1, 2, 3, 4}
	f(nums...)
}
