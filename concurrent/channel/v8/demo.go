package main

import "fmt"

// 向已关闭的chan 继续写入数据，会报错;
func main() {
	intChan := make(chan int)

	go func(cha chan int) {
		intChan <- 10
		close(intChan)
		// 向已关闭的chan 继续写入数据，会报错;
		intChan <- 20
	}(intChan)

	// 读取数据
	a := <-intChan
	fmt.Printf("接收数据: %v \n", a)
	b := <-intChan
	fmt.Printf("接收数据: %v \n", b)
	fmt.Println("程序运行结束!")

}

/** 输出
panic: send on closed channel

goroutine 6 [running]:
main.main.func1(0xc000016120, 0xc000016120)
	C:/Users/mark/go/src/go-study/concurrent/channel/v8/demo.go:13 +0x66
created by main.main
	C:/Users/mark/go/src/go-study/concurrent/channel/v8/demo.go:9 +0x76
*/
