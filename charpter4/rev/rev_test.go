package rev

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Start initializing.")
	result := m.Run() //运行go的测试，相当于调用main方法
	log.Println("Flush resources.")
	os.Exit(result) //退出程序
}

func TestRev1(m *testing.T) {
	fmt.Println("Start testing rev.go")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := []int{}
	fmt.Printf("Oring int slice is: %v\n", a)
	fmt.Printf("After rev,the slice is %v.\n", rev(a))
	fmt.Printf("Oring int slice is: %v\n", b)
	fmt.Printf("After rev,the slice is %v.\n", rev(b))
	fmt.Printf("Oring int slice is: %v\n", c)
	fmt.Printf("After rev,the slice is %v.\n", rev(c))
}

func TestRev2(m *testing.T) {
	fmt.Println("Start testing reverse slice 4 elements to the left.")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Origin slice is %v.\n", a)
	rev(a[:4])
	fmt.Printf("After rev,the slice is %v.\n", a)
	rev(a[4:])
	fmt.Printf("After rev,the slice is %v.\n", a)
	rev(a)
	fmt.Printf("After rev,the slice is %v.\n", a)
}

func TestRev3(m *testing.T) {
	fmt.Println("Start testing reverse with pointer.")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Origin slice is %v.\n", a)
	rev2(&a)
	fmt.Printf("After rev,the slice is %v.\n", a)
}

func TestSliceEqual(m *testing.T) {
	fmt.Println("Start testing func sliceEqual.")
	origin := []string{"a", "a", "a", "a", "A", "a"}
	a, b, c, d, e := origin[:3], origin[1:4], origin[2:5], origin[:4], origin[1:1]
	fmt.Printf("is %v and %v equal? Result is %v\n", a, b, sliceEqual(a, b))
	fmt.Printf("is %v and %v equal? Result is %v\n", a, c, sliceEqual(a, c))
	fmt.Printf("is %v and %v equal? Result is %v\n", a, d, sliceEqual(a, d))
	fmt.Printf("is %v and %v equal? Result is %v\n", a, e, sliceEqual(a, e))
}

func TestCleanDuplicate(m *testing.T) {
	fmt.Println("Start testing func cleanDuplicate.")
	origin1 := []string{"a", "a", "b", "b", "A", "c"}
	origin2 := []string{""}
	fmt.Printf("Clear the duplicate strings in %v: Result is %v\n", origin1, cleanDuplicate(origin1))
	fmt.Printf("Clear the duplicate strings in %v: Result is %v\n", origin2, cleanDuplicate(origin2))
}

func TestExtrudeSpace(m *testing.T) {
	fmt.Println("Start testing func cleanDuplicate.")
	b := []byte{33,22,44,00,00,56,71}
	fmt.Printf("Origin byte is %v: Result is %v\n", b, extrudeSpace(b))
}

func TestReverseNoSpace(m *testing.T) {
	fmt.Println("Start testing func cleanDuplicate.")
	b := []byte{33,22,44,00,00,56,71}
	fmt.Printf("Origin byte is %v: Result is %v\n", b, reverseSaveSpace(b))
}
