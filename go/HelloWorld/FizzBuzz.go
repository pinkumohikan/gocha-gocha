package main

import (
	"fmt"
	"strconv"
)

const FIZZ = 3
const BUZZ = 5
const MAX_NUMBER = 100

func judge(num int) string {
	if num%FIZZ == 0 && num%BUZZ == 0 {
		return "FizzBuzz"
	} else if num%FIZZ == 0 {
		return "Fizz"
	} else if num%BUZZ == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(num)
	}
}

func main() {
	for i := 1; i <= MAX_NUMBER; i++ {
		fmt.Println(judge(i))
	}
}
