package main

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestGoodHtml(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id="short"><span class="special">hi</span></p><br/>
</body>
</html>
`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}
	stdout = new(bytes.Buffer)
	goodHtml(doc)
	_, err = html.Parse(bytes.NewReader(stdout.(*bytes.Buffer).Bytes()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFunc(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id="short"><span class="special">hi</span></p><br/>
</body>
</html>
`
	stdin = strings.NewReader(input)
	stdout = new(bytes.Buffer)
	main()
	_, err := html.Parse(bytes.NewReader(stdout.(*bytes.Buffer).Bytes()))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
