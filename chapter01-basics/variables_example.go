package main

import "fmt"

// 运行这个文件: go run variables.go
func main() {
	// 1. 变量声明的不同方式

	// 方式1: var 关键字声明
	var message string
	message = "这是通过var声明的变量"
	fmt.Println(message)

	// 方式2: var 关键字声明并初始化
	var greeting string = "你好！"
	fmt.Println(greeting)

	// 方式3: 类型推断
	var number = 42
	fmt.Printf("数字: %d, 类型: %T\n", number, number)

	// 方式4: 短变量声明（最常用）
	name := "张三"
	isStudent := true
	height := 175.5

	fmt.Printf("姓名: %s\n", name)
	fmt.Printf("是学生: %t\n", isStudent)
	fmt.Printf("身高: %.1f cm\n", height)

	// 2. 多变量声明
	var a, b, c int = 1, 2, 3
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	x, y, z := "苹果", "香蕉", "橙子"
	fmt.Printf("水果: %s, %s, %s\n", x, y, z)

	// 3. 常量
	const pi = 3.14159
	const language = "Go"

	fmt.Printf("圆周率: %.5f\n", pi)
	fmt.Printf("编程语言: %s\n", language)

	// 4. 基本数据类型
	var (
		整数  int     = 100
		浮点数 float64 = 99.99
		布尔值 bool    = true
		字符串 string  = "Go语言真棒！"
		字节  byte    = 65  // A的ASCII码
		符文  rune    = '中' // Unicode字符
	)

	fmt.Printf("整数: %d\n", 整数)
	fmt.Printf("浮点数: %.2f\n", 浮点数)
	fmt.Printf("布尔值: %t\n", 布尔值)
	fmt.Printf("字符串: %s\n", 字符串)
	fmt.Printf("字节: %c (%d)\n", 字节, 字节)
	fmt.Printf("符文: %c (%d)\n", 符文, 符文)
}
