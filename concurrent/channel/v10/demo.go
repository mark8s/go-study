package main

import (
	"fmt"
	"time"
)

// 缓冲Channel
func main() {
	fmt.Printf("开始时间: %v \n", time.Now().Unix())
	// 创建一个缓冲区为2的整数型chan
	intChan := make(chan int, 2)
	intChan <- 100
	fmt.Printf("结束时间: %v \n", time.Now().Unix())
	fmt.Printf("intChan2  类型: %T 缓冲大小: %v \n", intChan, cap(intChan))
	fmt.Println("程序运行结束!")
}
