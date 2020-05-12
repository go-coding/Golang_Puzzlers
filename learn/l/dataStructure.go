package main

import (
	"fmt"
)

// go 数据结构

// go 基本类型 ( basic types )

// src/builtin/builtin.go

//var a int

//var m complex128 //

// bool
// string
// int int8 int16 int32 int64
// uint unit8 uint16 unit32 uint64
// byte // unit8 的别名
// rune // int32 的别名 代表一个Unicode码
// float32 float64

// complex64(实部虚部都是一个float32)和complex128 (实部虚部都是一个float64)。复数的形式为：re+im i。其中re为实部，im为虚部。
// 复数可以使用内置的complex()函数或者包含虚部数值的常量来创建。复数的各个部分可以使用内置函数real()和imag()函数获得。
// complex64 complex128 // 复数类型

// 标准库中有一个复数包 math/cmplx 提供复数各种通用的操作函数。

// 零值 （这个过程可以类比 JavaScript 变量申明）
// 变量在定义时没有明确的初始化时会赋值为_零值_。
// 零值是：
//
// 数值类型为 `0`，
// 布尔类型为 `false`，
// 字符串为 `""`（空字符串）。

// 单引号和双引号

//Go中，双引号是用来表示字符串string，其实质是一个byte类型的数组，单引号表示rune类型

//Go中字符串是一个不可变的值类型，内部用指针指向UTF-8字节数组。因此可以用索引号访问某字节，也可以用len()函数来获取字符串所占的字节长度。例如：

// func main() {
// 	var str string = "hello world"

// 	fmt.Println(str[3:5])

// 	for i := 0; i < 10; i++ {
// 		fmt.Printf(" i eq %v \n", i)
// 	}
// }

// 类型转换

// Go 的在不同类型之间的项目赋值时需要显式转换

// 类型推导  变量的类型由右值推导得出。

// 常量

// 常量的定义与变量类似，只不过使用 const 关键字。

// 常量可以是字符、字符串、布尔或数字类型的值。

// 常量不能使用 := 语法定义。

// 复杂类型 （struct、slice、map）

// struct
// 一个结构体（`struct`）就是一个字段的集合。
// 结构体字段使用点号来访问。
// 结构体字段可以通过结构体指针来访问。 通过指针间接的访问是透明的。
//

// 数组
// 类型 [n]T 是一个有 n 个类型为 T 的值的数组。 数组的长度是其类型的一部分，因此数组不能改变大小。

// slice
// 一个 slice 会指向一个序列的值，并且包含了长度信息。  []T 是一个元素类型为 T 的 slice。
// slice 由函数 make 创建。这会分配一个零长度的数组并且返回一个 slice 指向这个数组

// 数组和slice 区别

//声明赋值
//array
//
//
//arr := [6]int{2, 3, 5, 7, 11, 13}
//slice

//sli := arr[1:4]
//arr[1:4]就是对arr做的切片
//
//本质
//array就是array
//
//slice是array元素的引用，官方把slice叫做underlying array、Slices are like references to arrays
//
//从使用效果上看，slice完全就像是array元素的指针
//
//用函数传递例子更可以看出区别：slice传递的是地址，效果和指针相同，而array是复制元素
//
//
//package main
//
//import "fmt"
//
//func arrDo(x [6]int) {
//	x[2] = 333
//}
//
//func sliDo(x []int) {
//	x[2] = 333
//}
//
//func main() {
//	arr := [6]int{11, 22, 33, 44, 55, 66}
//	sli := []int{1, 2, 3, 4, 5, 6}
//
//	arrDo(arr)
//	sliDo(sli)
//
//	fmt.Println("arr:", arr)
//	fmt.Println("sli:", sli)
//}
//输出

//arr: [11 22 33 44 55 66]
//sli: [1 2 333 4 5 6]
//可以看出通过函数传递的array，在函数里把值修改后在函数外看不到。而slice是可以看到
//
//长度是否可变
//array是固定长度，占用内存中固定的一块地址，因此不可能通过删除array元素来导致长度变化
//
//slice是可变长度，内存地址可扩展
//
//2种声明方式
//array
//
//普通方式：

//显式声明数组arr，长度为6，元素类型为int
//var arr [6]int
//arr[0] = 2
//arr[1] = 3
//arr[2] = 5
//arr[3] = 7
//arr[4] = 11
//arr[5] = 13
//literal方式：

//arr := [6]int{2, 3, 5, 7, 11, 13}

//slice
//
//普通方式：

//显式声明切片sli，其中可见元素为arr的第2个元素地址到第4个元素地址，但容量为第2个元素地址到最后一个元素地址
//var sli []int = arr[1:4]
//literal方式：
//
//
//sli := []int{2, 3, 5, 7, 11, 13}

//类型
//array的类型示例：[6]int
//
//slice的类型示例：[]int

// range

// for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

// map
// map 映射键到值。
//
// map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
// map 的文法跟结构体文法相似，不过必须有键名。
//
//修改 map
//在 map m 中插入或修改一个元素：
//
//m[key] = elem
//获得元素：
//
//elem = m[key]
//删除元素：
//
//delete(m, key)
//通过双赋值检测某个键存在：
//
//elem, ok = m[key]
//如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
//
//同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。