package main

import "fmt"

func main() {

	// 创建一个整数型chan
	intChan := make(chan int)

	// 创建一个写入channel的协程
	go func(intChan chan int) {
		intChan <- 10
		// 关闭通道
		close(intChan)
	}(intChan)

	// 读取数据
	a := <-intChan
	fmt.Printf("接收数据: %v \n", a)
	b := <-intChan
	fmt.Printf("接收数据: %v \n", b)
	// 重复关闭(此处会报错)
	close(intChan)
	fmt.Println("程序运行结束!")

}
