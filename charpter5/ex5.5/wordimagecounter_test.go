package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"os"
	"testing"
)

func TestCounter(t *testing.T) {
	source, _ := ioutil.ReadFile("./file_xindong_com.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	wordsExpects := 42
	imagesExpects := 0
	words, images := Counter(doc)
	if words != wordsExpects {
		t.Errorf("Words count wrong. Expects: %d, got %d\n", wordsExpects, words)
	}

	if images != imagesExpects {
		t.Errorf("Images count wrong. Expects: %d, got %d\n", imagesExpects, images)
	}
}

func TestFunc(t *testing.T) {
	var res string
	os.Args = []string{"", "https://github.com/"}
	stdout = new(bytes.Buffer)
	main()
	ret := stdout.(*bytes.Buffer).String()
	if len(ret) == 0 {
		t.Errorf("No output")
	}

	os.Args = []string{"wiCounter"}
	stderr = new(bytes.Buffer)
	expects := "usage: wiCounter URL\n"
	main()
	res = stderr.(*bytes.Buffer).String()
	if res != expects {
		fmt.Printf("Error response is: %s", res)
	}
}
