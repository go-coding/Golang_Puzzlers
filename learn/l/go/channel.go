package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {

	//chs := make([]chan int, 10)
	//
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//
	//	log.Printf("chs ==>> %v ", chs[i])
	//
	//	go Count(chs[i])
	//}
	//
	//for _, ch := range (chs) {
	//	<-ch
	//}

	//ch2 := make(chan int)
	//
	//ch3 := <-chan int(ch2)
	//
	//ch4 := chan<- int(ch2)

	log.Println("cpu => ", runtime.NumCPU())

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i, &i)
		}()
	}

	time.Sleep(time.Second)

	//log.Printf("ch2 ==>> %v  ch2-T ==> %T", ch2, ch2)
	//log.Printf("ch3 ==>> %v  ch3-T ==> %T", ch3, ch3)
	//log.Printf("ch4 ==>> %v  ch4-T ==> %T", ch4, ch4)

}

func Count(ch chan int) {

	log.Printf("ch ==>> %v  ch-x ==> %x", ch, ch)

	ch <- 1
}
