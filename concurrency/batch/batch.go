package main

import (
	"fmt"
	"sync"
)

type S struct {
	Integer int64
	String  string
}

const sizeBatch = 10

func BatchProcessTest(data []int, m int) {
	ch := make(chan struct{}, sizeBatch)
	var wg sync.WaitGroup
	for _, i := range data {
		wg.Add(1)
		ch <- struct{}{}
		x := i
		go func() {
			defer wg.Done()
			d := x * m
			fmt.Println(d)
			<-ch
		}()
	}
	wg.Wait()
	fmt.Println("done processing all data")
}

func ProcessBatch(i int, r int) {
	data := make([]int, 0, r)
	for n := 0; n < r; n++ {
		data = append(data, n)
	}
	BatchProcessTest(data, i)
}
