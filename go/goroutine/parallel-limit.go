package main

import (
	"fmt"
	"time"
	"sync"
)

// refs: http://jxck.hatenablog.com/entry/20130414/1365960707
func main() {
	const ParallelLimit = 3
	const Max = 10

	limit := make(chan int, ParallelLimit)

	var wg sync.WaitGroup

	for i := 0; i < Max; i++ {
		wg.Add(1)
		limit <- 1

		go func(i int) {
			time.Sleep(time.Second)
			fmt.Println(fmt.Sprintf("%d done", i))
			<-limit
			wg.Done()
		}(i)
	}

	wg.Wait()
}
