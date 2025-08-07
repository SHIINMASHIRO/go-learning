// Package main 是Go程序的入口包
package main

// 导入fmt包用于格式化输出
import "fmt"

// main函数是程序的入口点
func main() {
	// 使用fmt.Println()输出文本
	fmt.Println("你好，世界！")
	fmt.Println("欢迎来到Go语言的世界！")

	// 声明变量
	var name string = "Go学习者"
	fmt.Printf("你好，%s！\n", name)

	// 短变量声明
	age := 25
	fmt.Printf("我今年%d岁了\n", age)
}
