package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

type c int

func add1(a int, b c) int {
	return a + int(b)
}

type inc interface {
	Inc() int
}

type Num struct {
	a int
	b string
	c string
}

type Cou struct {
	a string
	b int
	c int
}

func add2(a, b inc) int {
	return a.Inc() + b.Inc()
}

func (d Num) Inc() int {
	return d.a
}

func (e Cou) Inc() int {
	return e.b
}

func main() {
	var m Num = Num{
		a: 1,
	}
	var n Num = Num{
		a: 2,
	}
	var j Cou = Cou{
		b: 1,
	}
	var k Cou = Cou{
		a: "4",
	}
	// var j, k Cou
	x := add(1, 2)
	y := add1(1, 2)
	z := add2(m, n)
	w := add2(j, k)
	f := add2(m, j)
	fmt.Println(x, y, z, w, f)
}
