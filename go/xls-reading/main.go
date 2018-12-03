package main

import (
	"fmt"
	xls2 "github.com/extrame/xls"
)

func main() {
	xls, err := xls2.Open("list.xls", "utf-8")
	if err != nil {
		panic(err)
	}

	readCells := 100
	fmt.Println(xls.ReadAllCells(readCells))
}