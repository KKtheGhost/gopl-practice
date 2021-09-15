package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
	//"unicode"
	//"unicode/utf8"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	for k, v := range counts {
		fmt.Printf("%s == %d \n", k, v)
	}
}
