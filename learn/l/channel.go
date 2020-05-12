package main

import (
	"fmt"
	_ "sync"
)

// channel
// channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

//ch <- v    // 将 v 送入 channel ch。
//v := <-ch  // 从 ch 接收，并且赋值给 v。

// （“箭头”就是数据流的方向。）

// 和 map 与 slice 一样，channel 使用前必须创建：

// ch := make(chan int)

//channel支持3种类型（通过%T看到的）：
//
//读写类型：chan int
//
//只读类型：<-chan int，叫做receive-only
//
//只写类型：chan<- int，叫做send-only

// func sum(a []int, c chan int) {
// 	sum := 0
// 	for _, v := range a {
// 		sum += v
// 	}
// 	c <- sum // 将和送入 c
// }

// func main() {
// 	a := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)

// 	go sum(a[:len(a)/2], c)
// 	go sum(a[len(a)/2:], c)

// 	x, y := <-c, <-c // 从 c 中获取

// 	fmt.Println(x, y, x+y)
// }

// 缓冲 channel
// make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：
// ch := make(chan int, 100)

// 向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞。
