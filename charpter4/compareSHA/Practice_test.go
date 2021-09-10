package compareSHA

import (
	"crypto/sha256"
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

func TestCompareSHA(m *testing.T) {
	log.Println("Start func compareSHA Testing...")
	a, b := "x", "x"
	c := "X"
	res1 := compareSHA(sha256.Sum256([]byte(a)), sha256.Sum256([]byte(b)))
	fmt.Println(res1)
	res2 := compareSHA(sha256.Sum256([]byte(a)), sha256.Sum256([]byte(c)))
	fmt.Println(res2)
}

func TestCompareSHA2(m *testing.T) {
	log.Println("Start func compareSHA2 Testing...")
	a, b := "x", "x"
	c := "X"
	res1 := compareSHA2(sha256.Sum256([]byte(a)), sha256.Sum256([]byte(b)))
	fmt.Println(res1)
	res2 := compareSHA2(sha256.Sum256([]byte(a)), sha256.Sum256([]byte(c)))
	fmt.Println(res2)
}
