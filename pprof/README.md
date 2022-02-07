### pprof

#### 介绍
Go语言中的`pprof`指对于指标或特征的分析（Profiling），通过分析不仅可以查找到程序中的错误（内存泄漏、race冲突、协程泄漏），也能对程序进行优化（例如CPU利用率不足）。

由于Go语言运行时的指标不对外暴露，因此有标准库net/http/pprof和runtime/pprof用于与外界交互。其中net/http/pprof提供了一种通过http访问的便利方式，用于用户调试和获取样本特征数据。

#### 收集样本
在通过`pprof`进行特征分析时，需要执行两个步骤：收集样本和分析样本

pprof 采样数据主要有三种获取方式:

net/http/pprof: 通过 http 服务获取Profile采样文件，简单易用，适用于对应用程序的整体监控。底层也是通过 runtime/pprof 实现。

runtime/pprof: 手动调用runtime.StartCPUProfile或者runtime.StopCPUProfile等 API来生成和写入采样文件，灵活性更高。

go test: 通过 go test -cpuprofile cpu.pprof -memprofile mem.pprof生成采样文件,适用对函数进行针对性测试。其中-cpuprofile：生成CPU性能测试信息; -memprofile：生成内存占用信息；

