package main

import (
	"fmt"
)

func receive(ch1, ch2, ch3, quit chan int) {
	for i := 0; i < 4; i++ {
		fmt.Printf("receive %d from ch1\n", <-ch1)
		fmt.Printf("receive %d from ch2\n", <-ch2)
		fmt.Printf("receive %d from ch3\n", <-ch3)
	}
	quit <- 0
}

// select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。当多个都准备好的时候，会随机选择一个。

func send(ch1, ch2, ch3, quit chan int) {
	for i := 0; i < 10; i++ {
		select {
		case ch1 <- i:
			fmt.Printf("send %d to ch1\n", i)
		case ch2 <- i:
			fmt.Printf("send %d to ch2\n", i)
		case ch3 <- i:
			fmt.Printf("send %d to ch3\n", i)
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("default")
		}
	}
}

//当 select 中的其他条件分支都没有准备好的时候，`default` 分支会被执行。
//
//为了非阻塞的发送或者接收，可使用 default 分支：

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	quit := make(chan int)

	go receive(ch1, ch2, ch3, quit)

	send(ch1, ch2, ch3, quit)

}

//func m()  {
//	for i := 0; i < 10; i++ {
//		fmt.Println(<-c)
//	}
//	quit <- 0
//}
