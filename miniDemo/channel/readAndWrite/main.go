package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 10)

/*
 关闭chan 一般放在写入数据的地方；
 chan 一直可以读取数据，读取到的数据跟chan 关闭前是否存在元素有关，存在正常读取，第二个返回值 true, 反之 false;
*/

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("Write data: ", i)
		}
		close(ch)
		fmt.Println("close chan!")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Millisecond)
			v, ok := <-ch
			if ok {
				fmt.Println("read data: ", v)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Exit!")
}
