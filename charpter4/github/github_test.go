package github

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

func TestSearchIssues(m *testing.T) {
	log.Println("Start func searchIssues() Testing...")
	keywords := []string{"KKtheGhost", "xdanger"}
	res, err := SearchIssues(keywords)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
