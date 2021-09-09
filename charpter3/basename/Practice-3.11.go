package basename

import (
	"bytes"
	"strings"
)

// Func comma without recursion.
func comma3(s string) string {
	var buf bytes.Buffer
	startIndex := 0
	if s[0] == '+' || s[0] == '-' {
		// buf.WriteByte(s[0])
		startIndex = 1
	}
	endIndex := strings.Index(s, ".")
	if endIndex == -1 {
		endIndex = len(s)
	}

	realString := s[startIndex:endIndex]
	n := len(realString)
	if n <= 3 {
		buf.WriteString(realString)
	} else {
		for i := n; i > 0; i-- {
			if i%3 == 0 && i != n {
				buf.WriteByte(',')
			}
			buf.WriteByte(s[n-i])
		}
	}
	buf.WriteString(s[endIndex:])
	return buf.String()
}
