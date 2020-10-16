package core

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Test(ctx context.Context) {
	wg := new(sync.WaitGroup)
	fmt.Println("core.Test")

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("goroutine running : %d \n", i)
		}(ctx, i)
	}
	wg.Wait()
}
