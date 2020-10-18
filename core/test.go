package core

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Test is test
func Test(ctx context.Context) error {
	wg := new(sync.WaitGroup)
	fmt.Println("core.Test")

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(ctx context.Context, i int) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			select {
			case <-ctx.Done():
			default:
				fmt.Printf("goroutine running : %d \n", i)
				http.Get(fmt.Sprintf("http://localhost:8000/?i=%d", i))
			}
		}(ctx, i)
	}
	wg.Wait()
	return nil
}

// wg := new(sync.WaitGroup)
// fmt.Println("core.Test")

// for i := 1; i <= 10; i++ {
// 	wg.Add(1)
// 	go func(done chan bool, i int) {
// 		defer wg.Done()
// 		time.Sleep(1 * time.Second)
// 		fmt.Printf("goroutine running : %d \n", i)
// 	}(done, i)
// }
// wg.Wait()
