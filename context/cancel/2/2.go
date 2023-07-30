package main

import (
	"context"
	"fmt"
	"time"
)

func helper(ctx context.Context, tk *time.Ticker) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exiting from GO routine!")
			return
		case <-tk.C:
			fmt.Println("tick!")
		}
	}
}

func work(ctx context.Context) {
	ctx, cancelFn := context.WithCancel(ctx)

	go func() {
		time.Sleep(5 * time.Second)
		cancelFn()
	}()

	helper(ctx, time.NewTicker(time.Second))
}

func main() {
	work(context.Background())
}
