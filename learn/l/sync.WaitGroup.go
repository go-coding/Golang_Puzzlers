
package main

import (
	"log"
	"math/rand"
	"time"
	"sync"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		c := rand.Intn(5)
		d := time.Second * time.Duration(c) / 2
		log.Println("greeting => ", greeting , "d => ", d, "c => " ,c)
		time.Sleep(d)
	}
	wg.Done() // 通知当前任务已经完成。
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // 注册两个新任务。
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	wg.Wait() // 阻塞在这里，直到所有任务都已完成。
}

