package main

import "fmt"

func main(){
    
// 创建一个容量为4的channel 
ch := make(chan int, 4)
// 创建4个协程，作为生产者
for i := 0; i < 4; i++ {
  go func() {
    ch <- 7
  }()
}
// 创建4个协程，作为消费者
for i := 0; i < 4; i++ {
    go func() {
      o := <-ch
      fmt.Println("received:", o)
    }()
}
}
