package main

import "fmt"

// 从空接口中取值

type I interface {
}

func main() {

	num := 10
	var i I = num
	fmt.Printf("输出变量i: %v \n", i)

	var c int = i.(int)
	fmt.Printf("输出变量c: %v \n", c)
}
