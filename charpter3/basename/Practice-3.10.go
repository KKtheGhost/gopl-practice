package basename

import "bytes"

// Func comma without recursion.
func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i := n; i > 0; i-- {
		if i%3 == 0 && i != n {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[n-i])
	}
	return buf.String()
}
