### testing
测试


#### 介绍
testing 包为Go 语言提供自动化测试的支持。通过 go test 命令来执行单元测试文件，单元测试文件命名格式为: xxx_test.go,在单元测试文件中,根据测试类型不同可以分为:功能测试函数、基准测试函数,区别如下:

|类型 |函数格式 | 作用 |
| --- | --- | --- |
| 功能测试函数 | TestXXX(t *testing.T)| 测试函数功能是否正常 |
| 基准测试函数 | BenchmarkXXX(b *testing.B) | 测试函数的性能 |

#### 操作参考
> http://liuqh.icu/2021/07/07/go/package/24-testing/

