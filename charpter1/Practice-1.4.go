package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	K := 0
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countlines(f, counts)
			for _, n := range counts {
				if n > 1 {
					K++
				}
			}
			if K > 0 {
				fmt.Println(arg)
			}
			f.Close()
		}
	}
}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
