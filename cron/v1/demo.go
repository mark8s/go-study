package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

// 快速使用

func main() {
	// 开启秒字段支持
	c := cron.New(cron.WithSeconds())
	// 传统表达式写法: 每秒执行一次
	c.AddFunc("0/1 * * * * *", func() {
		fmt.Println("传统表达式: ", time.Now().Format("2006-01-02 15:04:05"))
	})

	// 预定义表达式
	c.AddFunc("@every 2s", func() {
		fmt.Println("预定义表达式: ", time.Now().Format("2006-01-02 15:04:05"))
	})
	// 启动
	c.Start()
	// 防止程序直接退出
	for {
	}

}
