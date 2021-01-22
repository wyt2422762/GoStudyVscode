package goroutine

//goroutine go 协程

import (
	"fmt"
	"time"
)

func helloWorld(a, b int) {
	fmt.Println("Hello World")
}

/* 参数为int类型的channel(通道) */
func helloWorldGoroutine(ch chan int) {
	fmt.Println("Hello World Goroutine")
	time.Sleep(5 * time.Second)
	ch <- 1 //向通道写入值1
}

func helloWorldDemo() {
	ch := make(chan int)
	/*
		在一个函数调用前加上go关键字，这次调用就会在一个新的goroutine(go协程)中并发执行。当被调用
		的函数返回时，这个goroutine也自动结束了。需要注意的是，如果这个函数有返回值，那么这个
		返回值会被丢弃。
	*/
	go helloWorldGoroutine(ch)
	<-ch //从channel(通道)中读数据，这里会阻塞
	time.Sleep(5 * time.Second)
	fmt.Println("Hello World Demo")
}


