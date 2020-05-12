# Go In Action

几个初步的 Go 语言特点

- 语言级并发
- 强大的标准库
- 垃圾回收
- 类型安全
- 简洁的面向对象

一共 25 个关键字：

```
package         包声明
import          包引入

const           常量

var             变量

type            类型

goto            流程控制
if              流程控制
else            流程控制
break           流程控制
continue        流程控制
default         流程控制
switch          流程控制
case            流程控制
fallthrough     流程控制
for    	        流程控制

func            函数相关
return          函数相关
defer           函数相关

struct          面向对象
interface       面向对象

map             数据结构
range           数据结构

go              并发相关
chan            并发相关
select          并发相关
```


## 包声明

```
package main

``` 

默认地将源码目录作为包的 package name。

之外：某个名称（包内方法或变量或常量或结构体等）在包外是否可见，就取决于其首个字符是否为大写字母

## 包引用

```
// 方式一
import "fmt"

// 方式二
import (
    "fmt"
    "log"
)
```


- 支持相对路径和绝对路径
- 绝对路径：import "lib/math" math.Sin
- 别名导入：import m "lib/math" m.Sin
- 点导入：import . "lib/math" Sin
- 下划线导入：import _ "lib/math" 只调用包中 init 函数


支持相对路径和绝对路径
绝对路径 import "xxx/yyy" 会引入 $GPATH/src/xxx/yyy
点导入：import . "fmt" 类似于 using namespace fmt;
别名导入: import f "fmt"
下划线导入：import _ "xxx/xxx/xxx" 只调用包中 init 函数


## 常量

## 变量

## 类型

## iota 枚举

## array

## slice

## map

## make,new 操作

## 零值

## 流程控制

## 函数

## 函数类型与传值
 
 
 
## panic 与 recover


## main 与 init 函数


## struct


## 面向对象



## 反射



## 并发

