package main

import (
	"errors"
	"fmt"
)

// 官方包

func main() {
	// 方式一: 使用errors包下的New()
	err := createError(1)
	printError(err)
	//  方式二: 使用fmt包下的Errorf()
	err2 := createError(2)
	printError(err2)
}
func printError(err error) {
	if err != nil {
		fmt.Printf("err==> %v | err.Error() ==> %v | 类型==> %T \n", err, err.Error(), err)
	}
}

// 创建error对象
func createError(way int) error {
	if way == 1 {
		// 方式一: 使用errors包下的New()
		return errors.New("方式一: 使用errors包下的New() ")
	} else if way == 2 {
		// 方式二: 使用fmt包下的Errorf()
		return fmt.Errorf("方式二: 使用fmt包下的Errorf(...) ---> ")
	}
	return nil
}
