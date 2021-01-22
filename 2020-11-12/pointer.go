package main

import "fmt"

func main() {
	fmt.Println("pointer")
	// & 取地址 * 取数据值

	var ip *int // 声明一个int类型的指针
	i := 2333   // 声明并赋值一个int类型的变量

	ip = &i // 指针指向变量地址

	fmt.Printf("i 变量的地址是: %x\n", &i)

	fmt.Printf("ip 变量储存的指针地址: %x\n", ip) // 指针变量存的地址

	fmt.Printf("*ip 变量的值: %d\n", *ip) // 使用指针访问值

}
