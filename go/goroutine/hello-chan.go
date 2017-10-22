package main

import (
	"fmt"
)

func passHoge(ch chan string) {
	ch<-"hoge"
}

func main() {
	ch := make(chan string)

	go passHoge(ch)

	fmt.Println(<-ch)
}
