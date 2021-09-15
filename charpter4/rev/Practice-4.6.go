package rev

import "unicode"

// 练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func extrudeSpace(b []byte) []byte {
	end := len(b)
	for i := 0; i < len(b); i++ {
		if !unicode.IsSpace(rune(b[i])) {
			continue
		} else {
			spaceCount := 0
			cpyStart := i
			for unicode.IsSpace(rune(b[i])) {
				i++
				spaceCount++
			}
			end = end - spaceCount + 1
			copy(b[cpyStart:], b[cpyStart+spaceCount-1:])
			i = cpyStart
		}
	}
	return b[:end]
}
