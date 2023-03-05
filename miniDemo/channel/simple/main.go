package main

func main() {
	// 没有缓冲区，同步读写情况下，没有接收者导致死锁
	QUse()
	// 设置缓冲区，数据写入缓冲区，chan 关闭后可以正常读取， 但是关闭后不能写入
	AUse1()
	// 并发读取，一边写，一边读
	AUse2()
}

func QUse() {
	ch := make(chan int)
	ch <- 1
	close(ch)
	<-ch
}

func AUse1() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	<-ch
}

func AUse2() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	close(ch)
	go func() {
		<-ch
	}()
}
