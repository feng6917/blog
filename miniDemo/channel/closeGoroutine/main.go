package main

import (
	"fmt"
	"time"
)

func main() {
	// 没啥区别，关闭chan就行

}

func A1() {
	var ch = make(chan int)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value == 1 {
					fmt.Println("Exit!")
					break
				}
				fmt.Println(value)
			}
		}
	}()

	ch <- 0
	ch <- 1
	close(ch)

	time.Sleep(1 * time.Second)
}

func A2() {
	var ch = make(chan int)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value%9 == 0 {
					fmt.Println("Exit!")
					break
				}
				fmt.Println(value)
			}
		}
	}()

	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	time.Sleep(1 * time.Second)
}

func A3() {
	var ch = make(chan int)
	var chS = make(chan bool)
	go func() {
		for {
			value, ok := <-ch
			if ok {
				if value%9 == 0 {
					fmt.Println("Exit!")
					chS <- true
					close(chS)
					break
				}
				fmt.Println(value)
			}
		}
	}()

	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		value := <-chS
		if value {
			close(ch)
		}
	}()

	time.Sleep(1 * time.Second)
}
