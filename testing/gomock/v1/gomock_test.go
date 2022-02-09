package v1

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBar(t *testing.T) {
	Convey("test mock", t, func() {
		// 创建gomock控制器，用来记录后续的操作信息
		controller := gomock.NewController(t)
		defer controller.Finish()
		// 调用mockgen生成的NewMockFoo方法
		m := NewMockFoo(controller)

		// 打桩代码 -开始
		m.EXPECT().
			DoSomething(). // 测试的方法
			Return(nil).   // 返回值
			Times(1)       // 调用次数
		// 打桩代码 -结束 当
		So(Bar(m), ShouldEqual, true)
		Convey("yes, when (return error)", func() {
			m.EXPECT().DoSomething().Return(errors.New("omg")).Times(1)
			fmt.Println("re: ", Bar(m))
			So(Bar(m), ShouldEqual, false)
		})
	})
}
