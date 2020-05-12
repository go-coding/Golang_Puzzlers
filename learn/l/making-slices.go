package main

import "fmt"

// func main() {
// 	a := make([]int, 5)
// 	printSlice("a", a)

// 	b := make([]int, 0, 5)
// 	printSlice("b", b)

// 	c := b[:2]
// 	printSlice("c", c)

// 	d := c[2:5]
// 	printSlice("d", d)
// }

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// len()可以用来查看数组或slice的长度
//
// cap()可以用来查看数组或slice的容量


// 在数组中由于长度固定不可变，因此len(arr)和cap(arr)的输出永远相同

//