package main

import (
	"fmt"
	"unsafe"
)

type Book struct {
	Age   int
	Title string
}

// 执行 golangci-lint run demo.go
func main() {
	b := Book{}
	fmt.Println("book size:", unsafe.Sizeof(b))
}
