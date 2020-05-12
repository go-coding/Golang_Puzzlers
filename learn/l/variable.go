package main

import (
	"log"
)

// go 变量申明

// 关键词 var 申明 变量

// 写法

// 一次申明一个变量
var foo int

// 一次申明多个变量
var bar, baz int

// 申明变量并初始化变量
var foo int = 1
var foo, bar int = 1, 2

// 短申明变量 （在函数中）

// := 简洁赋值语句可以替代 var 定义
// 函数外的每个语句都必须要关键词开始 var、 func 等

func t() {
	m := 1
	log.Printf("m eq %v", m)
}





