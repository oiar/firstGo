package main

import "fmt"

func main() {
	var b *int
	a := 10
	b = &a

	fmt.Printf("变量的地址: %x\n", &a)
	fmt.Printf("%d\n", a)
	fmt.Printf("%x\n", b)
	fmt.Printf("%x\n", &b)
	fmt.Printf("%d\n", *b)

	add(a)
	fmt.Println("调用了 add 之后 a 的值是 ", a)

	addPointer(&a)
	fmt.Println("调用了 addPointer 之后的值：", a)
}

func add(a int) {
	a++
	fmt.Println("我在 add 函数域里的值是", a)
}

func addPointer(a *int) {
	*a++
}
