package main

import "fmt"

// 阻塞特性的使用

func main() {

	// 创建一个整数型chan
	intChan := make(chan int)
	// 创建一个用于阻塞的chan
	boolChan := make(chan bool)

	// 创建一个写入数据的协程
	go func(cha chan int) {
		// 写入
		intChan <- 50
		fmt.Println("写入数据50")
		// 关闭通道
		close(intChan)
	}(intChan)

	// 创建一个读取数据的协程
	go func(intChan chan int, boolChan chan bool) {
		data, ok := <-intChan
		if ok {
			fmt.Printf("读取到数据 ->  %v \n", data)
			// 读取到数据后，给boolChan写入值
			boolChan <- true
			// 关闭用于的阻塞的chan
			close(boolChan)
		}
	}(intChan, boolChan)
	// 忽略接收，达到阻塞的效果。(如果不阻塞，则会直接输出: 程序运行结束!,不会等待协程执行)
	<-boolChan
	fmt.Println("程序运行结束!")

}
