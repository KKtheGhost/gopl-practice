package compareSHA

func compareSHA(sha1, sha2 [32]byte) bool {
	count := 0
	for i := 0; i < 32; i++ {
		if sha1[i] != sha2[i] {
			count++
			break
		}
	}
	return count == 0
}

func compareSHA2(sha1, sha2 [32]byte) bool {
	for i, _ := range sha1 {
		n, m := sha1[i], sha2[i]
		if n != m {
			return false
			break
		}
	}
	return true
}
