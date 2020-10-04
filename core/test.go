package core

import (
	"fmt"
	"sync"
	"time"
)

func Test() {
	wg := new(sync.WaitGroup)
	fmt.Println("core.Test")
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("goroutine running : %d \n", i)
		}(i)
	}
	wg.Wait()
}
