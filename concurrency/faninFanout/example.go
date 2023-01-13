package main

import (
	"context"
	"sync"

	"fmt"
	"time"
)

func FanInE(ctx context.Context, fetchers ...<-chan interface{}) <-chan interface{} {
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

func generate(data string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- data
			time.Sleep(1000)
		}
	}()

	return channel
}

type ProcessorE struct {
	jobChannel chan string
	done       chan *WorkerE
	workers    []*WorkerE
}
type WorkerE struct {
	nameE string
}

func (w *WorkerE) processJobE(data string, done chan *WorkerE) {
	// Use the data and process the job
	go func() {
		fmt.Println("Working on data ", data, w.nameE)
		time.Sleep(3000)
		done <- w
	}()

}

func GetProcessorE() *ProcessorE {
	p := &ProcessorE{
		jobChannel: make(chan string),
		workers:    make([]*WorkerE, 5),
		done:       make(chan *WorkerE),
	}
	for i := 0; i < 5; i++ {
		w := &WorkerE{nameE: fmt.Sprintf("<Worker - %d>", i)}
		p.workers[i] = w
	}
	p.startProcessE()
	return p
}

func (p *ProcessorE) startProcessE() {
	go func() {
		for {
			select {
			default:
				if len(p.workers) > 0 {
					w := p.workers[0]
					p.workers = p.workers[1:]
					w.processJobE(<-p.jobChannel, p.done)
				}
			case w := <-p.done:
				p.workers = append(p.workers, w)
			}
		}
	}()
}

func (p *ProcessorE) postJobE(jobs <-chan string) {
	p.jobChannel <- <-jobs
}

func E() {
	source := generate("data string")
	process := GetProcessorE()

	for i := 0; i < 12; i++ {
		process.postJobE(source)
	}

}
