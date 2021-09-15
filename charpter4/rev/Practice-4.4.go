package rev

// eg. a := []int{1,2,3,4,5,6,7}, x := 5
// length = 7
// tmp := make([]int, 5)
// tmp = []int{1,2,3,4,5}
// a = []int{6,7}
// a[2:]=[]int{1,2,3,4,5}
// a = []int{6,7,1,2,3,4,5}

func rotate(a []int, x int) {
	length := len(a)
	tmp := make([]int, x)
	copy(tmp, a[:x])
	copy(a, a[x:])
	copy(a[length-x:], tmp)
}
