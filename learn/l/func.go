package main

import "fmt"

// 函数
// 函数也是值。
// Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	//pos, neg := adder(), adder()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(-2*i),
	//	)
	//}
	//

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

func fibonacci() func() int {
	m, n := 0, 1

	//fmt.Printf(" m ==>> %v, n  ==>> %v", m, n)

	return func() int {
		t := m
		m, n = n, n+m

		return t
	}
}
