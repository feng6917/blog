package main

import (
	"context"
	"fmt"
	"time"
)

func f1(in chan struct{}) {

	time.Sleep(1 * time.Second)
	in <- struct{}{}

}

func f2(in chan struct{}) {
	time.Sleep(3 * time.Second)
	in <- struct{}{}
}

func main() {
	// 简单使用测试
	// DoCtxTime()

	/*
		场景：两个goroutine,第一个2秒执行完毕，第二个3秒执行完毕。
	*/

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	go func() {
		go f1(ch1)
		select {
		case <-ctx.Done():
			fmt.Println("f1 timeout")
			break
		case <-ch1:
			fmt.Println("f1 Done")
		}
	}()

	go func() {
		go f2(ch2)
		select {
		case <-ctx.Done():
			fmt.Println("f2 timeout")
			break
		case <-ch2:
			fmt.Println("f2 done")
		}
	}()

	time.Sleep(time.Second * 5)
}

func DoCtxTime() {
	tm, ctx := context.WithTimeout(context.Background(), 2*time.Second)

	defer ctx()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("3 s")
	case <-time.After(4 * time.Second):
		fmt.Println("4 s")
	case <-tm.Done():
		fmt.Println("Done")
	}
}
