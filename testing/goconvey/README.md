### GoConvey


#### 介绍
GoConvey是一个非常非常好用的Go测试框架，它直接与go test集成，提供了很多丰富的断言函数，能够在终端输出可读的彩色测试结果，并且还支持全自动的Web UI。

#### 下载
```gotemplate
go get github.com/smartystreets/goconvey
```

#### 操作参考
> https://www.liwenzhou.com/posts/Go/golang-unit-test-5/

#### 断言方法

```yaml
So(thing1, ShouldEqual, thing2)
So(thing1, ShouldNotEqual, thing2)
So(thing1, ShouldResemble, thing2)		// 用于数组、切片、map和结构体相等
So(thing1, ShouldNotResemble, thing2)
So(thing1, ShouldPointTo, thing2)
So(thing1, ShouldNotPointTo, thing2)
So(thing1, ShouldBeNil)
So(thing1, ShouldNotBeNil)
So(thing1, ShouldBeTrue)
So(thing1, ShouldBeFalse)
So(thing1, ShouldBeZeroValue)

```

#### WebUI
goconvey提供全自动的WebUI，只需要在项目目录下执行以下命令。
```yaml
goconvey
```
默认就会在本机的8080端口提供WebUI界面，十分清晰地展现当前项目的单元测试数据。

