package main

import (
	"fmt"
	"sync"
	"time"
)

func job(index int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("index: %d\n", index)
}

var pool chan struct{}

func main() {
	maxNum := 10
	pool = make(chan struct{}, maxNum)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		pool <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-pool
			}()
			job(i)
		}(i)
	}
	wg.Wait()
}
