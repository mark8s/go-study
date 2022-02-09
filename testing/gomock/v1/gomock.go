package v1

import "fmt"

type Foo interface {
	DoSomething() error
}

// 调用方在调用被调用方时，必须先定义接口（go语言的接口由使用方定义），这样才能进行mock

func Bar(f Foo) bool {
	err := f.DoSomething()
	if err != nil {
		// do something, 需要测试该代码
		fmt.Println("error")
		return false
	}
	// do something, 需要测试该代码
	fmt.Println("ok")
	return true
}

// 生成mock代码文件
//mockgen -source=gomock.go -destination=mock_gomock.go -package=mock
//参数说明：
//-source: 包含`interface`定义的源文件
//-destination: 生成的mock代码输出到哪里
//-package: 生成的mock代码文件的包名
