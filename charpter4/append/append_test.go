package append

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

func TestAppendInt1(m *testing.T) {
	log.Println("Start func appendInt1 Testing...")
	var x []int
	x = appendInt(x, 6)
	fmt.Println(x)
	x = appendInt(x, 99)
	fmt.Println(x)
}

func TestAppendInt2(m *testing.T) {
	log.Println("Start func appendInt1 Testing...")
	var x []int
	x = appendInt2(x, 6, 66, 666)
	fmt.Println(x)
	x = appendInt2(x, x...)
	fmt.Println(x)
}
