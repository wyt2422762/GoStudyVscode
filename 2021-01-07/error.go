package main

//错误处理 重点：error(接口) defer(延迟执行) panic(类似于java中的throw) recove(类似于java中的catch)

//defer 延迟执行，类似于java中的finally, 会在方法支持完毕前或者异常退出前执行，defer方法会被压入栈中，多个defer方法按先进后出的顺序执行

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//简单示例

func readFile(srcPath string) {
	srcFile, err := os.Open(srcPath) //打开文件，默认为只读权限
	if err != nil {
		fmt.Println("打开源文件失败")
	}

	defer srcFile.Close() //延迟执行 源文件关闭操作

	//读取源文件
	br := bufio.NewReader(srcFile)

	for {
		line, isPrefix, err1 := br.ReadLine() //按行读取文件
		if err1 != nil {                      //读取文件异常(如果异常为EOF，表示文件读取完成)
			fmt.Println("读取源文件错误")
			break
		}
		fmt.Printf("isPrefix : %v\n", isPrefix)
		fmt.Println(string(line))
	}
}

func copyFile(srcPath string, destPath string) {
	srcFile, err := os.Open(srcPath) //打开文件，默认为只读权限
	if err != nil {
		fmt.Println("打开源文件失败")
	}

	defer srcFile.Close() //延迟执行 源文件关闭操作

	destFile, err1 := os.OpenFile(destPath, os.O_WRONLY, 0666) //以写的方式打开问价
	if err1 != nil {
		fmt.Println("打开目标文件失败")
	}

	defer destFile.Close() //延迟执行 目标文件关闭操作

	bytesWritten, err2 := io.Copy(destFile, srcFile) //拷贝文件内容
	if err2 != nil {
		fmt.Println("拷贝文件失败")
	}

	fmt.Printf("拷贝的字节数：%d\n", bytesWritten)

	err3 := destFile.Sync() // 将文件内容flush到硬盘中
	if err3 != nil {
		fmt.Println("写入文件失败")
	}

	fmt.Printf("完成")

}

func throw(){
	panic("抛出的异常") //抛出一个异常
}


func main() {
	//readFile("E:/测试.txt")
	//copyFile("E:/测试.txt", "E:/123.txt")

	defer func(){
		if r := recover(); r != nil{ //捕获异常，相当于java中的catch
			fmt.Printf("捕获异常：%v\n", r)
		}
	}()

	throw()

}
