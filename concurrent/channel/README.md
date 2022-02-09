### channel

#### 介绍

channel即Go的通道，是协程之间的通信机制。一个channel是一条通信管道，它可以让一个协程通过它给另一个协程发送数据。每个channel都需要指定数据类型，即channel可发送数据的类型。Go语言主张通过数据传递来实现共享内存，而不是通过共享内存来实现数据传递。

##### 创建语法

```go
// 声明方式1
var cha1 chan 数据类型
cha1 = make(chan 数据类型)

// 声明方式2
cha1 := make(chan 数据类型)

```

##### 发送数据

通过channel发送数据需要使用特殊的操作符＜-，需要注意的是: channel发送的值的类型必须与channel的元素类型一致。

##### 接收数据

```go
// 方式一: ch 指的是通道变量
data := <- ch
//方式二: data 表示接收到的数据。未接收到数据时，data为channel类型的零值，ok（布尔类型）表示是否接收到数据
data, ok := <- ch

```

执行该语句时channel将会阻塞，直到接收到数据并赋值给data变量。

##### Channel的阻塞特性

- channel默认是阻塞的。
- 当数据被发送到channel时会发生阻塞，直到有其他Goroutine从该channel中读取数据。
- 当从channel读取数据时，读取也会被阻塞，直到其他Goroutine将数据写入该channel。

使用 `<- boolChan` 阻塞主线程

##### 关闭channel

发送方写入完毕后需要主动关闭channel，用于通知接收方数据传递完毕。接收方通过data,ok := <- ch判断channel是否关闭，如果ok=false，则表示channel已经被关闭。

往关闭的channel中写入数据会报错：panic: send on closed channel。导致程序崩溃。

重复关闭chan，会崩溃

#### 缓冲Channel

默认创建的都是非缓冲channel，读写都是即时阻塞。缓冲channel自带一块缓冲区，可以暂时存储数据，如果缓冲区满了，就会发生阻塞。缓冲通道在发送时无需等待接收方接收即可完成发送过程，并且不会发生阻塞，只有当缓冲区满时才会发生阻塞。同理，如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞。

##### 语法

```go
// 声明 n:代表缓冲区大小
cha1 := make(chan T, n)
```





#### 操作参考

> http://liuqh.icu/2020/08/31/go/basic/17-channel/ 

