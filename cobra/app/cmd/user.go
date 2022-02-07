/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// userCmd 父命令
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "用户操作",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("用户列表: ", list)
	},
}

// userCmd 子命令（添加用户）
var userAddCmd = &cobra.Command{
	Use:   "add",
	Short: "添加用户; user add --name=?",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("添加用户：", name)
	},
}

var (
	name string
	list []string
)

// userCmd 子命令（删除用户）
var userDelCmd = &cobra.Command{
	Use:   "del",
	Short: "删除用户; user del --name=?",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("删除用户：", name)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)

	// 添加子命令到父命令
	userCmd.AddCommand(userAddCmd)
	userCmd.AddCommand(userDelCmd)

	// 添加参数
	// 持久标志(PersistentFlags): 当前命令接收，当前命令行和其所有子命令都有该参数选项
	userCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "用户名")
	// 本地标志(Flags): 仅仅当前命令有该参数选项。
	userCmd.Flags().StringSliceVarP(&list, "list", "l", []string{}, "用户列表")
}
