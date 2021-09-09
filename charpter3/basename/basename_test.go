// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
package basename

import (
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
func TestBasename1(t *testing.T) {
	log.Println("Start Testing func basename1")
	b1 := basename1("/a/b/c/d.go")
	log.Printf("basename1: result is '%v'\n", b1)
	b2 := basename1("/a/b/c/d.e.f.go")
	log.Printf("basename1: result is '%v'\n", b2)
}
func TestBasename2(t *testing.T) {
	log.Println("Start Testing func basename2")
	b1 := basename2("/a/b/c/d.go")
	log.Printf("basename1: result is '%v'\n", b1)
	b2 := basename2("/a/b/c/d.e.f.go")
	log.Printf("basename1: result is '%v'\n", b2)
}

func TestComma(t *testing.T) {
	log.Println("Start Testing func comma")
	c1 := comma("21321867489326")
	log.Printf("basename1: result is '%v'\n", c1)
	c2 := comma("243")
	log.Printf("basename1: result is '%v'\n", c2)
}

func TestComma2(t *testing.T) {
	log.Println("Start Testing func comma2")
	c1 := comma2("21321867489326")
	log.Printf("basename1: result is '%v'\n", c1)
	c2 := comma2("243")
	log.Printf("basename1: result is '%v'\n", c2)
}

func TestIntToString(t *testing.T) {
	log.Println("Start Testing func printints")
	c1 := intsToString([]int{1, 2, 3, 4, 5})
	log.Printf("basename1: result is '%v'\n", c1)
	c2 := intsToString([]int{134, 21, 6})
	log.Printf("basename1: result is '%v'\n", c2)
}

func TestComma3(t *testing.T) {
	log.Println("Start Testing func comma3")
	c1 := comma3("-21321867489326")
	log.Printf("basename1: result is '%v'\n", c1)
	c2 := comma3("243.3348")
	log.Printf("basename1: result is '%v'\n", c2)
	c3 := comma3("15")
	log.Printf("basename1: result is '%v'\n", c3)
	c4 := comma3("+90789348.22")
	log.Printf("basename1: result is '%v'\n", c4)
	c5 := comma3("0")
	log.Printf("basename1: result is '%v'\n", c5)
}

func TestIsDiff(t *testing.T) {
	log.Println("Start Testing func IsDiff")
	c1 := isDiff("acdf", "fdac")
	log.Printf("basename1: result is '%v'\n", c1)
	c2 := isDiff("", "")
	log.Printf("basename1: result is '%v'\n", c2)
	c3 := isDiff("", "21321")
	log.Printf("basename1: result is '%v'\n", c3)
}
