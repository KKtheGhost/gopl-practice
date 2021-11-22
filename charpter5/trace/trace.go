package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
    defer trace("bigSlowOperation")() // 函数结束时调用 trace 函数
    // ... do other slow work ...
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
    start := time.Now()
    log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}