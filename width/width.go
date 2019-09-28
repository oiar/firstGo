package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type S struct {
		a uint16
		b uint32
		c uint16
	}
	var s S
	fmt.Println(unsafe.Sizeof(s)) //12, why?

	var d [3]uint32
	fmt.Println(unsafe.Sizeof(d))

	e := struct{}{}
	f := struct{}{}
	fmt.Println(e == f)
}
