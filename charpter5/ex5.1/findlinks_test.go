package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"reflect"

	"golang.org/x/net/html"
)

func visitLoop(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitLoop(links, c)
	}
	return links
}

func TestFindLinks(t *testing.T) {
	source, _ := ioutil.ReadFile("./Free eBooks _ Project Gutenberg.html")
	data := bytes.NewBuffer(source)
	doc, _ := html.Parse(data)
	ret := visit(nil, doc)
	expects := visitLoop(nil, doc)

	if !reflect.DeepEqual(ret, expects) {
		t.Errorf("Results:%q\nExpects:%q", ret, expects)
	}
}