package main

import "fmt"

//发送方写入完毕后需要主动关闭channel，用于通知接收方数据传递完毕。接收方通过data,ok := <- ch判断channel是否关闭，如果ok=false，则表示channel已经被关闭。

func main() {
	// 创建一个整数型chan
	intChan := make(chan int)
	// 创建一个写入channel的协程
	go func(intChan chan int) {
		intChan <- 10
		intChan <- 20
		// 关闭通道
		close(intChan)
	}(intChan)

	// 读取数据
	a := <-intChan
	fmt.Printf("接收数据: %v \n", a)
	b := <-intChan
	fmt.Printf("接收数据: %v \n", b)

	// 此时的Chan已经关闭，而且里面的数据也都已经取完
	c := <-intChan
	fmt.Printf("接收数据: %v \n", c)
	fmt.Println("程序运行结束!")
}
