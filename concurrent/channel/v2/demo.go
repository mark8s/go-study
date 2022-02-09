package main

import (
	"fmt"
)

func main() {
	// 创建chan
	intChan := make(chan int)
	// 写入数据
	go func() {
		// 写入数据（协程写入）
		intChan <- 5
		fmt.Println("协程写入数据: 5")
	}()

	// 接收数据（主线程读取）
	data, ok := <-intChan
	if ok {
		fmt.Printf("主线程接收数据：%v \n", data)
		fmt.Println("chan 是协程之间的通信, 所以必须以goroutine的方式传入数据")
	}

}
