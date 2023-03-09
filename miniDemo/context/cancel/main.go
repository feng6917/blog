package main

import (
	"context"
	"fmt"
	"time"
)

/*
context可以用来在goroutine之间传递上下文信息，相同的context可以传递给运行在不同goroutine中的函数，
上下文对于多个goroutine同时使用是安全的，context包定义了上下文类型，可以使用background、TODO创建一个上下文，
在函数调用链之间传播context，也可以使用WithDeadline、WithTimeout、WithCancel 或 WithValue 创建的修改副本替换它，
听起来有点绕，其实总结起就是一句话：context的作用就是在不同的goroutine之间同步请求特定的数据、取消信号以及处理请求的截止日期。
*/

// context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来。
// context.TODO 应该只在不确定应该使用哪种上下文时使用；

/*
background它通常由主函数、初始化和测试使用，并作为传入请求的顶级上下文；
TODO是当不清楚要使用哪个 Context 或尚不可用时，代码应使用 context.TODO，后续在在进行替换掉，归根结底就是语义不同而已。
*/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("done!")
				return
			default:
				fmt.Println(time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
