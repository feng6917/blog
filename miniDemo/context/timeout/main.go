package main

import (
	"context"
	"fmt"
	"time"
)

/*
withTimeout、WithDeadline不同在于WithTimeout将持续时间作为参数输入而不是时间对象，
这两个方法使用哪个都是一样的，看业务场景和个人习惯了，因为本质withTimout内部也是调用的WithDeadline。

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

*/

func main() {
	// 自动结束
	//do()

	// 手动结束
	do2()

}

func newContextWithTimeOut() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*3)
}

func deal(ctx context.Context) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("val: ", i)
		}
	}
}

func deal2(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("val: ", i)
			cancel()
		}
	}
}

func do() {
	ctx, cancel := newContextWithTimeOut()
	defer cancel()
	deal(ctx)
}

func do2() {
	ctx, cancel := newContextWithTimeOut()
	defer cancel()
	deal2(ctx, cancel)
}
