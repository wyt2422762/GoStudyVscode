package socket

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func simpleGetTest() {
	url := "http://www.baidu.com"
	resp, err := http.Get(url) //get方式请求url
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		fmt.Printf("get请求链接%s出错\n", url)
	}
	fmt.Println("返回状态码：", resp.StatusCode)
	//fmt.Println("返回结果：", resp.Body)
	io.Copy(os.Stdout, resp.Body) //将返回的body赋值到系统的输出流里
}

func simplePostTest() {
	url := "http://www.baidu.com"
	data := `{"code": "123"}`
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		fmt.Printf("post请求链接%s出错\n", url)
	}
	fmt.Println("返回状态码：", resp.StatusCode)
}

func simpleServerTest() {
	http.HandleFunc("/test", helloHandler2) //设置处理方法，类似于配置一个action
	http.HandleFunc("/", helloHandler1)     //设置处理方法，类似于配置一个action
	http.ListenAndServe(":8181", nil)       //监听本机端口8181
}

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r.Body.Close()
	}()
	name := r.URL.Query().Get("name") //获取get请求中的name参数
	fmt.Printf("helloHandler1: 请求的name参数：%s\n", name)
	io.WriteString(w, "Hello, "+name)
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r.Body.Close()
	}()
	name := r.URL.Query().Get("name") //获取get请求中的name参数
	fmt.Printf("helloHandler2: 请求的name参数：%s\n", name)
	io.WriteString(w, "Hello, "+name)
}
