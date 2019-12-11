package main

import (
	"fmt"

	"github.com/GinMu/golang-learning/goroutine"
)

func main() {
	goroutine.UnBufferedChannel()

	ch := make(chan int, 3)
	// 查看当前通道的大小
	fmt.Println(len(ch))
	// 发送3个整型元素到通道
	ch <- 1
	ch <- 2
	ch <- 3
	// 查看当前通道的大小
	fmt.Println(len(ch))
}
