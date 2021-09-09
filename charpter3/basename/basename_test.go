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
