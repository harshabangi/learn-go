package main

import (
	"context"
	"fmt"
	"time"
)

func helper(ctx context.Context, tk *time.Ticker, ch chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			ch <- struct{}{}
			close(ch)
			fmt.Println("exiting from GO routine!")
			return
		case <-tk.C:
			fmt.Println("tick!")
		}
	}
}

func work(ctx context.Context) {
	ctx, cancelFn := context.WithCancel(ctx)

	ch := make(chan struct{})
	tk := time.NewTicker(time.Second)

	go helper(ctx, tk, ch)
	time.Sleep(5 * time.Second)

	// call cancel
	cancelFn()

	<-ch
}

func main() {
	work(context.Background())
}
