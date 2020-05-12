package main

import "log"
import "testing"
import "time"

func Test_Go(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func() {
			log.Printf("i ==>> %v", i)
		}()
	}

	time.Sleep(100 * time.Millisecond)

	//for i := 0; i < 5; i++ {
	//	time.Sleep(100 * time.Millisecond)
	//	fmt.Println(s)
	//}

}
