package main

import "fmt"

func main() {
	var ch = make(chan int, 2)
	ch <- 1
	ch <- 2
	//do(ch)
	do2(ch)
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
