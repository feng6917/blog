package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 1)

func main() {
	/*
		场景：并发模拟吃饭场景，分为三个步骤，1，洗手，2，做饭，3，吃饭，依次进行。
	*/
	//A1()

	// 单个chan实现, 通过判断内部值
	//one.One()

}

func A1() {
	// 多个chan实现
	var chXS = make(chan struct{}, 1)
	var chZF = make(chan struct{}, 1)
	var chCF = make(chan struct{}, 1)

	chXS <- struct{}{}
	close(chXS)

	go func() {
		_, ok := <-chCF
		if ok {
			fmt.Println("开始做饭，准备吃饭！")
			fmt.Println("END!")
		}
	}()

	go func() {
		_, ok := <-chZF
		if ok {
			fmt.Println("开始做饭，准备吃饭！")
			chCF <- struct{}{}
			close(chCF)
		}
	}()

	go func() {
		_, ok := <-chXS
		if ok {
			fmt.Println("开始洗手，准备做饭！")
			chZF <- struct{}{}
			close(chZF)
		}
	}()

	time.Sleep(1 * time.Second)
}
