package main

import (
	"charpter2/popcount"
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, err := strconv.Atoi(os.Args[1:][0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "get1bit: %v\n", err)
		os.Exit(1)
	}
	res := popcount.PopCount(uint64(a))
	fmt.Println(res)
}
