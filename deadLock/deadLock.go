package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

//import (
//	"fmt"
//	"runtime"
//)
//
//type struct1 struct {
//	a, b int64
//	c, d float64
//	e *struct2
//}
//
//type struct2 struct {
//	f, g int64
//	h, i float64
//}
//
//func main() {
//	s1 := allocStruct1()
//	s2 := allocStruct2()
//
//	func () {
//		_ = allocStruct2()
//	}()
//
//	runtime.GC()
//
//	fmt.Printf("s1 = %X, s2 = %X\n", &s1, &s2)
//}
//
////go:noinline
//func allocStruct1() *struct1 {
//	return &struct1{
//		e: allocStruct2(),
//	}
//}
//
////go:noinline
//func allocStruct2() *struct2 {
//	return &struct2{}
//}

func main() {

	// Create server connection
	natsConnection, _ := nats.Connect("nats://yourusername:yoursecret@localhost:4222")
	log.Println("Connected")

	// Subscribe to subject
	log.Printf("Subscribing to subject 'bucketevents'\n")
	natsConnection.Subscribe("bucketevents", func(msg *nats.Msg) {

		// Handle the message
		log.Printf("Received message '%s\n", string(msg.Data)+"'")
	})

	// Keep the connection alive
	runtime.Goexit()
}