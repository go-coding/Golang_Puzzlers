package main

import "fmt"

func main() {

	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Printf("the time \n")
	//	}()
	//}


	select {} // fatal error: all goroutines are asleep - deadlock!



	ch := make(chan int, 10)

	for m := 1; m <= 10; m++ {
		ch <- m
	}

	Loop: for {
		select {
		case i := <-ch:
			fmt.Printf("select case i => %v \n", i)
		default:
			fmt.Printf("select default case  \n")
			break Loop
		}
	}

}
