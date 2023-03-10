package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 2)
	ch <- 1
	ch <- 2
	//do(ch)
	//do2(ch)

	do3(ch)
}

func do(ch chan int) {
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("val: ", v)
				if v == 2 {
					break
				}
			}
		}
	}
	fmt.Println("exit!")
}

// for select 无法直接使用break 关闭, 可以使用 break+label 关闭
func do2(ch chan int) {
Exit:
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Println("val: ", v)
				if v == 2 {
					break Exit
				}
			}
		}
	}
	fmt.Println("exit!")
}

// for select 可以使用chan关闭
func do3(ch chan int) {
	chStatus := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-ch:
				if ok {
					fmt.Println("val: ", v)
					if v == 2 {
						break
					}
				}
			// 使用time.After 仅为判断条件
			case <-time.After(10 * time.Millisecond):
				chStatus <- true
				close(chStatus)
				break
			}
		}
	}()
	<-chStatus
	fmt.Println("exit!")
}
