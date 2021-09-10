package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var t = flag.Int("t", 2, "-t 2|3|5")

func main() {
	flag.Parse()
	var stdin = make([]byte, 1024)
	a, err := os.Stdin.Read(stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	switch *t {
	case 3:
		fmt.Println(sha512.Sum384(stdin[0:a]))
	case 5:
		fmt.Println(sha512.Sum512(stdin[0:a]))
	default:
		fmt.Println(sha256.Sum256(stdin[0:a]))
	}
}
