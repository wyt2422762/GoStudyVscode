package main

import (
	"fmt"
)

type cusInt int //定义一个类型cusInt

func (a cusInt) cusFuncPlus(b cusInt) cusInt { //为cusInt类型添加一个方法
	return a - b
}

func (a cusInt) cusString(b string) string { //为cusInt类型添加一个方法
	return b
}

type cat struct { //定义一个结构体 猫
	age   int    //年龄
	color string //颜色
}

func main() {
	/* var a cusInt
	a = 1 */
	a := cusInt(1) //同上方注释的等价

	var b cusInt = 2
	c := a.cusFuncPlus(b)
	fmt.Printf("a = %d, c= %d\n", a, c)

	//Go语言中大多都为值传递
	var i int = 1
	var j int = i
	j++
	fmt.Printf("i: %d, j: %d\n", i, j)

	//数组也为值传递
	var ints [3]int = [3]int{1, 2, 3}
	var ints1 = ints
	ints[1]++
	fmt.Printf("ints值: %v, ints地址: %x, ints1值: %v, ints1地址: %x\n", ints, &ints[0], ints1, &ints1[0])

	//引用传递需要使用指针
	var ints2 [4]int = [4]int{1, 2, 3, 4}
	var ints3 = &ints2 //ints3 指向 ints2的地址  ints3是一个 *[4]int类型(指针类型)
	ints3[0]++

	fmt.Printf("ints2: %v, ints3: %v, ints3类型: %T\n", ints2, ints3, ints3)

	var testi []int = []int{1, 2, 3}
	testIntArray(testi)
	fmt.Printf("testi值: %v, testi地址: %x\n", testi, &testi[0])

	var ii int = 1
	testInt(ii)
	fmt.Printf("ii值: %d, ii地址: %x\n", ii, &ii)

}

func testIntArray(i []int){
	fmt.Printf("方法中i的类型: %T, i的地址: %x\n", i, &i[0])
	i[0]++
}

func testInt(i int){
	fmt.Printf("方法中i的类型: %T, i的地址: %x\n", i, &i)
	i++
}


