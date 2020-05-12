package main

// 指针

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	fmt.Println(p)  // read i through the pointer
	fmt.Println(i)  // read i through the pointer
	fmt.Println(&i) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

// & 符号会生成一个指向其作用对象的指针。

// * 符号表示指针指向的底层的值。

// 这也就是通常所说的“间接引用”或“非直接引用”。
