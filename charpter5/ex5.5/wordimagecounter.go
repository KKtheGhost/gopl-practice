package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var stderr io.Writer = os.Stderr
var stdout io.Writer = os.Stdout

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("Parsing HTML: %s", err)
		return
	}
	words, images = Counter(doc)
	return
}

func Counter(n *html.Node) (words, images int) {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[len(u)-1]
		u = u[:len(u)-1]
		switch n.Type {
		case html.TextNode:
			if n.Parent.Data != "script" && n.Parent.Data != "style" {
				words += WordCount(n.Data)
			}
		case html.ElementNode:
			if n.Data == "img" {
				images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
	}
	return
}

func WordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(stderr, "Usage: %s URL\n", os.Args[0])
		return
	}
	url := os.Args[1]
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(stdout, "Words: %d\nImages: %d\n", words, images)
}
