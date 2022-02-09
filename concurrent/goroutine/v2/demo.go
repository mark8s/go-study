package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("张三")
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("李四")
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("王五")
	// 睡眠1秒
	time.Sleep(1 * time.Second)
	fmt.Println("运行结束")

}
