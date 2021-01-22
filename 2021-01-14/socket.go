package socket

import (
	"fmt"
	"net"
)

func simpleTest() {
	_, err := net.Dial("ip4:icmp", "www.baidu.com")
	if err != nil {
		fmt.Println("连接失败")
	}
}

func socketServerTest() {
	l, err := net.Listen("tcp", ":8181") //监听本机的8181端口
	if err != nil {
		fmt.Println("监听端口出现错误：", err)
	}
	fmt.Println("已成功监听8181端口")
	for {
		conn, err := l.Accept() //接收连接
		if err != nil {
			fmt.Println("接收连接出现错误：", err)
		}
		if conn != nil {
			fmt.Println("接收到来自客户端的连接")
			handleConnection(conn)
		}
	}
}

//处理连接的方法
func handleConnection(conn net.Conn){
	defer func ()  {
		fmt.Println("服务端关闭连接")
		conn.Close()
	}()
	//具体的处理方法
	buffer := make([]byte, 1024) //声明并初始化了一个slice(切片)，初始大小为1024
	n, err := conn.Read(buffer)
	if err != nil{
		fmt.Println("读取数据错误：", err)
	}
	fmt.Println("读取的字节数：", n)
	fmt.Println("读取的数据：", string(buffer))
}

func socketClientTest(){
	conn, err := net.Dial("tcp", "127.0.0.1:8181")
	if err != nil{
		fmt.Println("连接服务器失败：", err)
	}
	s := "测试"
	var data []byte = []byte(s)
	n, err := conn.Write(data)
	if err != nil{
		fmt.Println("向服务器写数据失败：", err)
	}
	fmt.Println("向服务器写的字节数：", n)
}
