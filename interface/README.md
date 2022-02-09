### 接口

#### 介绍
在Go语言中，接口是一组方法签名。接口指定了类型应该具有的方法，类型决定了如何实现这些方法。当某个类型为接口中的所有方法提供了具体的实现细节时，这个类型就被称为实现了该接口。接口定义了一组方法，如果某个对象实现了该接口的所有方法，则此对象就实现了该接口。

空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从空接口中取出原值。
#### 操作参考
> http://liuqh.icu/2020/08/28/go/basic/14-interface/
