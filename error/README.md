### error

#### 介绍

错误是指程序中出现不正常的情况，从而导致程序无法正常运行。Go语言中没有try...catch来捕获错误，而是通过defer+recover+panic模式来实现捕捉错误信息。

panic让当前的程序进入恐慌，中断程序的执行。panic()是一个内建函数，可以中断原有的控制流程。其功能类似于java中的throw。

Go语言中没有try...catch来捕获错误，一旦触发panic就会导致程序崩溃。在Go中是通过recover让程序恢复。值的注意的是recover()必须在延迟函数(defer)中有效。

#### 操作参考

> http://liuqh.icu/2020/08/29/go/basic/15-error/


