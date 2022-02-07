package os

import (
	"fmt"
	"os"
	"testing"
)

// 系统相关
func TestSystemInfo(t *testing.T) {
	hostname, _ := os.Hostname()
	// 主机名: LAPTOP-VH57ARI1
	fmt.Printf("主机名:%v \n", hostname)
	// 调用者所在进程的进程ID: 7208
	fmt.Printf("调用者所在进程的进程ID: %v \n", os.Getpid())
	// 调用者所在进程的进程的父进程ID: 10400
	fmt.Printf("调用者所在进程的进程的父进程ID: %v \n", os.Getppid())
}
