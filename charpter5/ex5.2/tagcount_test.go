package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestTagCount(t *testing.T) {
	source, _ := ioutil.ReadFile("./Free eBooks _ Project Gutenberg.html")
	stdin = bytes.NewBuffer(source)
	stdout = new(bytes.Buffer)
	main()
}