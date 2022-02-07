/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动http服务,使用方法: app server?`",
	Run: func(cmd *cobra.Command, args []string) {
		// 使用配置
		if appConfig.App.Port == "" {
			fmt.Println("port不能为空!")
			os.Exit(-1)
		}
		engine := gin.Default()
		_ = engine.Run(":" + appConfig.App.Port)
	},
}

var port string

func init() {
	rootCmd.AddCommand(serverCmd)
	// 接收参数port
	serverCmd.Flags().StringVar(&port, "port", "", "端口号")
}
