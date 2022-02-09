package main

import "fmt"

// 空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。

// 定义一个空接口

type A interface {
}

func main() {
	// 声明变量
	var a A
	// 保存整型
	a = 10
	fmt.Printf("保存整型: %v \n", a)
	// 保存字符串
	a = "hello word"
	fmt.Printf("保存字符串: %v \n", a)
	// 保存数组
	a = [3]float32{1.0, 2.0, 3.0}
	fmt.Printf("保存数组: %v \n", a)
	// 保存切片
	a = []string{"您", "好"}
	fmt.Printf("保存切片: %v \n", a)
	// 保存Map
	a = map[string]int{
		"张三": 22,
		"李四": 25,
	}
	fmt.Printf("保存map: %v \n", a)
	// 保存结构体
	a = struct {
		name string
		age  int
	}{"王麻子", 40}
	fmt.Printf("保存结构体: %v \n", a)

	// 声明一个空接口切片
	var aa []A
	// 保存任意类型数据到切片中
	aa = append(aa, 23, []string{"php", "go"}, map[string]int{"a": 1, "b": 2}, struct {
		city, province string
	}{"长沙", "湖南"})
	fmt.Printf("空接口切片: %v \n", aa)

}
