package core

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Test() {
	wg := new(sync.WaitGroup)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	fmt.Println("core.Test")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				fmt.Println("Done:", ctx.Err())
				return
			case r := <-time.After(2 * time.Second):
				fmt.Println("", r)
				fmt.Printf("goroutine running : %d \n", i)
			}
		}(ctx, i)
	}
	wg.Wait()
	cancel()
}
