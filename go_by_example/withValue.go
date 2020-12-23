package main

import (
	"context"
	"fmt"
	"time"
)

type Options struct {
	Interval time.Duration
}

func study(ctx context.Context, workName string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop", workName)
			return
		default:
			fmt.Println(workName, "send request")
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	vctx := context.WithValue(ctx, "name", &Options{})
	go study(vctx, "work1")
	go study(vctx, "work2")

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}
