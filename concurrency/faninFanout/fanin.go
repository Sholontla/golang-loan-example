package main

import (
	"context"
	"sync"
)

func FanIn(ctx context.Context, fetchers ...<-chan interface{}) <-chan interface{} {
	combinedFetcher := make(chan interface{})
	// 1
	var wg sync.WaitGroup
	wg.Add(len(fetchers))

	// 2
	for _, f := range fetchers {
		f := f
		go func() {
			// 3

			defer wg.Done()
			for {
				select {
				case res := <-f:
					combinedFetcher <- res
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	// 4
	// Channel cleanup
	go func() {
		wg.Wait()
		close(combinedFetcher)
	}()
	return combinedFetcher
}
