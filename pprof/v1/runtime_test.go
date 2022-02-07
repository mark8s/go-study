package main

// 使用runtime/pprof

import (
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func TestRuntimePProf(t *testing.T) {
	// 打开文件
	f, err := os.Create("./out.pprof")
	if err != nil {
		t.Errorf("文件打开失败:%v", err)
		return
	}
	// 调用
	err = pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	if err != nil {
		t.Errorf("StartCPUProfile:%v", err)
	}
	// 测试单独函数
	testPprof()
}

// 模拟内存使用增加
func testPprof() {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 200)
		}
		ch <- true
	}()
	<-ch
}

// 运行
// go test runtime_test.go -v
// go tool pprof -http=:8080 out.pprof
