package main

import "fmt"

func main() {
	var m = make(map[string]string)
	m["a"] = "red"
	m["b"] = "blue"
	fmt.Println(m, "what is the m ?")
	fmt.Println(len(m), "what is the len?")

	delete(m, "a")
	fmt.Println(m, "delete the a")
}
