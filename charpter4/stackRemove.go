package main

import "fmt"

// remove slice[2] with order
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 5, 6, 7, 8, 9
// i = 2
// slice[i:] = []int{7,8,9}
// slice[i+1:] = []int{8,9}
// slice = []int{5,6,8,9,nil}
// return slice[:len(slice)-1]=[]int{5,6,8,9}

// remove slice[2] without order, use last element to overwrite slice[2]
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	t := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2)) // "[5 6 8 9]"
	fmt.Println(remove2(t, 2))
}
