package main

import (
	"fmt"
	"time"
)

// NewTimer()创建一个新的计时器，它会在至少持续时间d之后将当前时间发送到其channel上。
func main() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Printf("开始时间 %v \n", time.Now())

	out := <-timer.C
	fmt.Printf("变量out->  类型: %T 值:%v  \n", out, out)
	fmt.Printf("开始时间 %v \n", time.Now())

}
