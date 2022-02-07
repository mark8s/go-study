package go_funk

import (
	"fmt"
	"github.com/thoas/go-funk"
	"testing"
)

func TestIndexOf(t *testing.T) {
	strArr := []string{"go", "java", "c", "java"}
	// 具体类型
	fmt.Println("c: ", funk.IndexOfString(strArr, "c"))
	// 验证第一次出现位置
	fmt.Println("java: ", funk.IndexOfString(strArr, "java"))
	// 任意类型
	fmt.Println("go: ", funk.IndexOf(strArr, "go"))
	// 不存在时返回-1
	fmt.Println("php: ", funk.IndexOfString(strArr, "php"))
}

// 遍历切片
func TestForeachSlice(t *testing.T) {
	// 从左边遍历
	var leftRes []int
	funk.ForEach([]int{1, 2, 3, 4}, func(x int) {
		leftRes = append(leftRes, x*2)
	})
	fmt.Println("ForEach:", leftRes)
	// 从右边遍历
	var rightRes []int
	funk.ForEachRight([]int{1, 2, 3, 4}, func(x int) {
		rightRes = append(rightRes, x*2)
	})
	fmt.Println("ForEachRight:", rightRes)
}

// 判断是否为空
func TestIsEmpty(t *testing.T) {
	// 空结构体
	fmt.Println("空结构体", funk.IsEmpty([]int{}))
	// 空字符串
	fmt.Println("空字符串:", funk.IsEmpty(""))
	// 判断数字0
	fmt.Println("0:", funk.IsEmpty(0))
	// 判断字符串'0'
	fmt.Println("'0':", funk.IsEmpty("0"))
	// nil
	fmt.Println("nil:", funk.IsEmpty(nil))
}
