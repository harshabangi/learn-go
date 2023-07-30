package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func work(i int, ctx context.Context, wg *sync.WaitGroup, fn func()) {
	defer wg.Done()

	for {
		select {
		case <-time.After(time.Duration(2*i) * time.Second):
			fn()
			fmt.Printf("exiting from Go Routine: %d\n", i)
			return
		case <-ctx.Done():
			fmt.Printf("exiting from Go Routine: %d\n", i)
			return
		}
	}
}

func doMain(ctx context.Context, cancel context.CancelFunc) {

	var (
		wg sync.WaitGroup
		fn = func() {}
	)

	for i := 1; i <= 5; i++ {
		i := i

		if i == 3 {
			fn = func() { cancel() }
		}
		wg.Add(1)
		go work(i, ctx, &wg, fn)
	}

	wg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	doMain(ctx, cancel)
}
