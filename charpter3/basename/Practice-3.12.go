package basename

import "strings"

// Func isDiff in every byte between s1 & s2.
// return False means no difference
// return True means has difference
func isDiff(s1, s2 string) bool {
	pick := 0
	for i := 0; i < len(s1); i++ {
		if strings.Count(s1, s1[i:i+1]) == strings.Count(s2, s1[i:i+1]) {
			pick++
		}
	}
	if !strings.EqualFold(s1, s2) && pick == len(s1) && pick == len(s2) {
		return true
	}
	return false
}
