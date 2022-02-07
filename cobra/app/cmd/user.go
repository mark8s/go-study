/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"unicode/utf8"
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
	// 位置参数限制
	//Args:  cobra.RangeArgs(1, 3),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("参数数量不对")
		}
		// 判断姓名长度
		count := utf8.RuneCountInString(args[0])
		fmt.Printf("%v %v \n", args[0], count)
		if count > 4 {
			return errors.New("姓名长度过长")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("添加用户：", args[0])
		fmt.Println("位置参数(args):", args)
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

	// 设置参数必需
	/*err := userAddCmd.MarkFlagRequired("name")
	if err != nil {
		fmt.Println("--name 不能为空")
		return
	}*/

}
