package main

import "fmt"

type People struct {
}

func main() {
	// 创建整数型无缓冲chan
	intChan := make(chan int)
	fmt.Printf("intChan类型：%T ,值：%v \n", intChan, intChan)

	// 创建一个空接口类型chan，可以存放任意类型数据
	interfaceChan := make(chan interface{})
	fmt.Printf("interfaceChan类型：%T ,值：%v \n", interfaceChan, interfaceChan)

	// 创建一个指针chan
	peopleChan := make(chan *People)
	fmt.Printf("peopleChan类型：%T ,值：%v \n", peopleChan, peopleChan)
}

// 输出
//intChan类型：chan int ,值：0xc000016120
//interfaceChan类型：chan interface {} ,值：0xc000016180
//peopleChan类型：chan *main.People ,值：0xc0000161e0
