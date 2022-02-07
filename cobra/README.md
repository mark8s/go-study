### cobra
项目地址：https://github.com/spf13/cobra

#### 介绍
cobra是一个用来构建现代CLI工具的库。相比flag标准库，它提供更多方便的特性和功能。Cobra 由 Go 项目成员和 hugo 作者 spf13 创建，已经被许多流行的 Go 项目采用，比如 GitHub CLI 和 Docker CLI。

#### 特性概览
- 使用cobra add cmdname可快速的创建子命令cli
- 全局、局部和级联的标志
- 自动生成commands和flags的帮助信息
- 自动识别 -h、--help 等帮助标识
- 支持自定义帮助信息，用法等的灵活性。
- 可与viper 紧密集成

#### 安装
```go
go get -u github.com/spf13/cobra/cobra
```
#### 参考
> http://liuqh.icu/2021/11/07/go/package/28-cobra/

#### 一般步骤
- 创建app目录
- 进入后执行 `cobra init`
- 添加go mod，执行 `go mod init`、`go mod tidy`
- 为app 命令添加一个子命令server，执行`cobra add server`
- 在app根目录执行`go build .`
- 执行app命令，如`app server -p 8080`
- 添加子命令 `userCmd.AddCommand(userAddCmd)`
- 添加参数 `userCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "用户名")`


#### 参数限制
- NoArgs : 如果有任何位置参数，该命令将报告错误。

- MinimumNArgs(int) :至少传 N 个位置参数，否则报错。

- ArbitraryArgs: 接受任意个位置参数。

- MaximumNArgs(int) : 最多传N 个位置参数，否则报错。

- ExactArgs(int) : 传入位置参数个数等于N，否则报错。

- RangeArgs(min, max) : 传入位置参数个数 min<= N <= max，否则报错


















