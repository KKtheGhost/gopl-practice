package rev

// 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func cleanDuplicate(s []string) []string {
	l := len(s)
	var res []string
	if l == 0 {
		return s
	}
	res = append(res, s[0])
	for i := 0; i < l-1; i++ {
		if s[i] != s[i+1] {
			res = append(res, s[i+1])
		}
	}
	return res
}
