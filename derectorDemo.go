package main

import (
	"fmt"
	"time"
)

// 为函数类型设置别名提高代码可读性
type MultPlyFunc func(int, int) int

// 乘法运算函数
func multiply(a,b int) int  {
	return a * b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间

func exeTime(f MultPlyFunc) MultPlyFunc  {
	return func(a int, b int) int {
		start := time.Now()
		c := f(a, b) // 执行乘法运算函数
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c
	}
}

func main()  {
	a := 2
	b := 8
	// 通过修饰器调用乘法函数，返回的是一个匿名函数
	decorator := exeTime(multiply)
	c := decorator(a,b)
	fmt.Printf("%d x %d = %d\n", a, b, c)
}