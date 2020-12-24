package main

import (
	"context"
	"fmt"
	"time"
)

func work1(ctx context.Context, workName string) {
	for {
		select {
		case <- ctx.Done():
			fmt.Println("stop", workName)
			return
		default:
			fmt.Println("send request", workName)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go work1(ctx, "work1")
	go work1(ctx, "work2")
	time.Sleep(3 * time.Second)

	cancel()
	time.Sleep(3 * time.Second)
}