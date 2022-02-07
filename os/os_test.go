package os

import (
	"fmt"
	"os"
	"os/exec"
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

// 程序退出
func TestExit(t *testing.T) {
	fmt.Println("调用前打印...")
	// 调用退出程序：code范围应在 0 <= x <= 125
	os.Exit(0)
	// 后面代码不会执行
	fmt.Println("调用后，这里不会输出")
}

// 环境变量相关
func TestEnv(t *testing.T) {
	// 系统所有环境变量
	fmt.Printf("%+v \n", os.Environ())
	// 设置环境变量
	_ = os.Setenv("my-name", "mark8s")
	// 获取环境变量
	fmt.Printf("获取环境变量: %v \n", os.Getenv("my-name"))
	// 清空所有环境变量
	os.Clearenv()
	fmt.Printf("清空环境变量后:%+v \n", os.Environ())
}

// 在环境变量PATH中搜索可执行文件
func TestLookPath(t *testing.T) {
	path, err := exec.LookPath("go")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(path)
}

// 使用Run()执行命令
func TestRun(t *testing.T) {
	cmd := exec.Command("sleep", "3s")
	// 具体执行
	err := cmd.Run()
	if err != nil {
		t.Errorf("执行失败:%v\n", err)
	}

	fmt.Printf("cmd.Path: %v \n", cmd.Path)
	fmt.Printf("cmd.Args: %v \n", cmd.Args)
}
