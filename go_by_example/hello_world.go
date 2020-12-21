package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("hello, world!")
	//
	//for i := 0; i < 6; i++ {
	//	if i % 2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}
	//
	//i := 2
	//fmt.Print("write", i)
	//switch i {
	//case 1:
	//	fmt.Println("one")
	//case 2:
	//	fmt.Println("two")
	//case 3:
	//	fmt.Println("three")
	//}
	//
	//t := time.Now()
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("it's before noon")
	//default:
	//	fmt.Println("it's after noon")
	//}
	//
	//whatIam:= func(i interface{}) {
	//	switch t := i.(type) {
	//	case string:
	//		fmt.Println(t, "string")
	//	case int:
	//		fmt.Println(t, "int")
	//	case bool:
	//		fmt.Println(t, "bool")
	//	}
	//}
	//
	//whatIam(true)
	//whatIam(1)
	//whatIam("hey")
	//
	//c()
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("Recover in main from error:", err)
	//	}
	//	fmt.Println("Execution finished")
	//}()
	//
	//go func() {
	//	panic("panic in goroutine")
	//}()
	defer func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recover from error:", err)
			}
		}()
		if err := recover(); err != nil {
			fmt.Println("Recover from error:", err)
			panic("again")
		}
	}()

	panic("whatever")

	time.Sleep(1 * time.Second)
}

func c() (a int) {
	defer func() {
		a++
		fmt.Println(a, "a")
	}()
	return 1
}
