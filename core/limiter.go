package core

import (
	"context"
	"log"
	"sync"
	"time"
)

var threads = 2

// Limiter is Limiter
func Limiter(ctx context.Context) error {
	endpoints := []string{"/", "/ok", "/.git", "/fsecure", "/non", "//"}

	wg := new(sync.WaitGroup)
	jobs := make(chan string)

	// initialisation des workers
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case url, ok := <-jobs:
					if !ok { // no more jobs
						return
					}
					// if an error occurs while handling this job,
					// dont return, just break
					time.Sleep(2 * time.Second)
					log.Println(i, " : ", url)
				}
			}
		}(i)
	}

	// iteration sur les donnees à traiter
	for _, endpoint := range endpoints {
		select {
		case <-ctx.Done():
			break
		case jobs <- endpoint: // bloquant
		}
	}
	close(jobs)
	wg.Wait()

	return nil
}
