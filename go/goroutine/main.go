package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go printNumber(i)
	}

	// printNumberが実行されるまで待つ
	time.Sleep(time.Second)
}

func printNumber(num int) {
	fmt.Println(num)
}
