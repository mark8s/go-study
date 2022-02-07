package main

import "fmt"

// %v  只输出所有的值
// %+v 先输出字段名字，再输出该字段的值
// %#v 先输出结构体名字值，再输出结构体（字段名字+字段的值）*/

type student struct {
	id   int32
	name string
}

func main() {
	s := &student{
		id:   1,
		name: "mark",
	}

	fmt.Printf("a=%v	\n", s)
	fmt.Printf("a=%+v	\n", s)
	fmt.Printf("a=%#v	\n", s)

	// 输出
	/*
		a=&{1 mark}
		a=&{id:1 name:mark}
		a=&main.student{id:1, name:"mark"}
	*/
}
