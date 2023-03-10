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
	// 正常终止
	//a1()
	// --------
	// 父级终止 子级不终止    父级终止 子级终止
	//a2()
	// 子级终止 父级终止     子级终止结束，父级继续运行，任务完成终止结束
	//a3()
	// 子级终止 父级不终止   子级终止结束，父级不受影响
	//a4()

	// 验证a2 是否跟在父级函数创建子函数有关，验证结果，无关系，父级终止，子级即终止
	a5()

}

func a1() {
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

func a2() {
	var ctx, nctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		nctx, _ = context.WithCancel(ctx)
		go func(ctx context.Context) {
			for {
				time.Sleep(time.Second)
				select {
				case <-ctx.Done():
					fmt.Println("child done!")
					return
				default:
					fmt.Println("child: ", time.Second)
				}
			}
		}(nctx)
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	// 终止
	cancel()
	time.Sleep(1 * time.Second)
	// 终止
	//ncancel()
	//time.Sleep(1 * time.Second)
}

func a3() {
	var ctx, nctx context.Context
	var cancel, ncancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		nctx, ncancel = context.WithCancel(ctx)
		go func(ctx context.Context) {
			for {
				time.Sleep(time.Second)
				select {
				case <-ctx.Done():
					fmt.Println("child done!")
					return
				default:
					fmt.Println("child: ", time.Second)
					// 子级终止
					ncancel()
				}
			}
		}(nctx)
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	// 父级终止
	cancel()
	time.Sleep(1 * time.Second)
}

func a4() {
	var ctx, nctx context.Context
	var ncancel context.CancelFunc
	ctx, _ = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		nctx, ncancel = context.WithCancel(ctx)
		go func(ctx context.Context) {
			for {
				time.Sleep(time.Second)
				select {
				case <-ctx.Done():
					fmt.Println("child done!")
					return
				default:
					fmt.Println("child: ", time.Second)
					// 子级终止
					ncancel()
				}
			}
		}(nctx)
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)
	time.Sleep(5 * time.Second)
	// 父级终止
	// cancel()
	time.Sleep(1 * time.Second)
}

func a5() {
	var ctx, nctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("parent done!")
				return
			default:
				fmt.Println("parent: ", time.Second)
			}
		}

	}(ctx)

	nctx, _ = context.WithCancel(ctx)
	go func(ctx context.Context) {
		for {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("child done!")
				return
			default:
				fmt.Println("child: ", time.Second)
			}
		}
	}(nctx)
	time.Sleep(5 * time.Second)
	// 终止
	cancel()
	time.Sleep(1 * time.Second)

	time.Sleep(5 * time.Second)
}
