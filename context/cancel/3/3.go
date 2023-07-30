package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func helper(ctx context.Context, tk *time.Ticker, id string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("exiting from GO routine: %s\n", id)
			return
		case <-tk.C:
			fmt.Printf("tick!: %s\n", id)
		}
	}
}

func work(ctx context.Context) {
	ctx, cancelFn := context.WithCancel(ctx)

	var wg sync.WaitGroup
	wg.Add(1)

	for i := 0; i < 2; i++ {
		i := i
		wg.Add(1)

		go func(ctx context.Context, i int, wg *sync.WaitGroup) {
			defer wg.Done()

			for j := 0; j < 2; j++ {
				j := j
				wg.Add(1)

				go func(ctx context.Context, j int, wg *sync.WaitGroup) {
					defer wg.Done()
					helper(ctx, time.NewTicker(1*time.Second), fmt.Sprintf("%d%d", i, j))
				}(ctx, j, wg)
			}

			helper(ctx, time.NewTicker(1*time.Second), fmt.Sprintf("%d", i))
		}(ctx, i, &wg)
	}

	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		cancelFn()
	}()

	wg.Wait()
}

func main() {
	work(context.Background())
}
