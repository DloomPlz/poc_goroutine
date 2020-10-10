package core

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Test(ctx context.Context) {
	wg := new(sync.WaitGroup)
	fmt.Println("core.Test")
	c := make(chan os.Signal)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			go func() {
				<-c
				ctx.Done()
			}()
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
}
