package main

import "log"

/**
 ++、--是语句，不是表达式， *p++等同于(*p)++。

运算符有5层优先级：

Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=>  >=
    2             &&
    1             ||


https://www.php.net/manual/zh/language.operators.precedence.php

和 PHP js 均不一样

 */

func main()  {

	res := 1<<31 - 1

	b := 1<<7-1

	log.Printf("res ==>> %v b ==>> %v", res, b)
}

