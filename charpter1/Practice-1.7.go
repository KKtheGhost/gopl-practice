package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	wt := bufio.NewWriter(os.Stdout)
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		n, err := io.Copy(wt, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: write %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\n\n=================\nTotal byte size is: %v\n", n)
		wt.Flush()
	}
}
