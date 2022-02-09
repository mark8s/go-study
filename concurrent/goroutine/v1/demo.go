package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go echoNum("张三")
	go echoNum("李四")
	go echoNum("王五")

	time.Sleep(time.Second)
	fmt.Printf("func main end")
}

func echoNum(who string) {
	for i := 1; i < 3; i++ {
		fmt.Println(who + " " + strconv.Itoa(i))
	}
}
