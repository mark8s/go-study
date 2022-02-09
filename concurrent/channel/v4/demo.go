package main

import "fmt"

//循环接收数据，需要配合使用关闭channel，借助普通for循环和for ... range语句循环接收多个元素。
//遍历channel，遍历的结果就是接收到的数据，数据类型就是channel的数据类型。
//普通for循环接收channel数据，需要有break循环的条件；for … range会自动判断出channel已关闭，而无须通过判断来终止循环。

// 使用for i , 通过判断channel中是否有值来判断channel是否已经关闭

func main() {

	intChan := make(chan int)

	// 协程写入数据
	go func(cha chan int) {
		// 写入
		for i := 1; i < 5; i++ {
			intChan <- i
			fmt.Printf("写入数据 -> %v \n", i)
		}
		close(intChan)
	}(intChan)

	// 主线程接收数据
	for {
		out, ok := <-intChan
		// 判断通道是否关闭
		// 如果通道关闭，则ok为false，也就是说明channel中没有值了
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("接收数据 ==> %v \n", out)
	}
	fmt.Println("程序运行结束!")
}
