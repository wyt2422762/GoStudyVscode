package main

import (
	"fmt"
	//"github.com/wyt2422762/2020-12-09/cus"
)

func main() {
	/* res, _ := calc(1, 2)
	fmt.Println(res)
	cus.Test()

	cusFunc1(123, "str1", "str2")

	sum, sub := cusFunc2(3, 2)
	fmt.Printf("sum = %d, sub = %d\n", sum, sub)

	//匿名函数
	f := func(a int, b int) (sum int, sub int) {
		return a + b, a - b
	}
	f(3, 1)
	//直接调用匿名函数
	bb, cc := func(a int, b int) (sum int, sub int) {
		return a + b, a - b
	}(5, 9) //大括号后面跟参数表示直接调用函数
	fmt.Printf("sum = %d, sub = %d\n", bb, cc) */


	//闭包
	bibao := func(j int)(func()){
		return func(){
			j++
			fmt.Printf("闭包内的参数j = %d\n", j)
		}
	}

	bb := bibao(1)
	bb()
	bb()
	bb()
}

/* 结构 func 方法名(参数名 参数类型...)(返回值名 返回值类型...) */
func calc(a int, b int) (res int, err error) {
	c := a + b
	return c, nil
}

/* 不定参数, 不定参数必须是最后一个参数 */
func cusFunc1(a int, b ...string) {
	fmt.Printf("定参数a = %d\n", a)
	for i, v := range b {
		fmt.Printf("不定参数b[%d] = %s\n", i, v)
	}
}

/* 多返回值 */
func cusFunc2(a int, b int) (sum int, sub int) {
	return a + b, a - b
}
