package v3

import "fmt"

func main() {
	// 定义一个匿名延迟函数
	defer func() {
		err := recover()
		msg := fmt.Sprintf("err信息: %v", err)
		if err != nil {
			// 程序触发panic时，会被这里捕获
			fmt.Println(msg)
		}
	}()
	// 定义一个数组
	testPanic("请求失败")
}

// 故意抛出panic
func testPanic(err string) {
	panic("错误信息" + err)
}

/** 输出
err信息: 错误信息请求失败
*/
