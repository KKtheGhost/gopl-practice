package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
	"sort"
)

var stdout io.Writer = os.Stdout
var stdin io.Reader = os.Stdin

func main() {
	countTable, err := tagCount(stdin)
	if err != nil {
		log.Fatal(err)
	}
	var tags []string
	for tag := range countTable {
		tags = append(tags,tag)
	}
	sort.Strings(tags)
	for _,tag := range tags {
		fmt.Fprintf(os.Stdout, "%4d %s\n", countTable[tag],tag)
	}
}

func tagCount(r io.Reader) (map[string]int, error) {
	countTable := make(map[string]int, 0)
	t := html.NewTokenizer(r)
	var err error
	for {
		tag := t.Next()
		if tag == html.ErrorToken {
			break
		}
		name,ok := t.TagName()
		if ok {
			countTable[string(name)]++
		}
	}
	if err != io.EOF {
		return countTable,err
	}
	return countTable,nil
}