package main

import "fmt"

type cat struct { //定义一个结构体 猫
	age   int    //年龄
	color string //颜色
	strs  []int
}

func (m *cat) miao(str string) { //给cat结构体增加一个方法
	fmt.Printf("Miao: %s", str)
}

func (m *cat) changeColor(color string) {
	m.color = color
}

func (m cat) changeStrs1(strs []int) {
	m.strs = strs
}

func (m *cat) changeStrs2(strs []int) {
	m.strs = strs
}

func (m cat) addStrs1(){
	m.strs[0]++
}

func (m *cat) addStrs2(){
	m.strs[0]++
}

func main() {
	//var c *cat = new(cat)
	var c *cat = &cat{2, "yellow", []int{1}}
	fmt.Printf("c的类型为: %T\n", c)
	c.miao("123")
	c.changeColor("red")
	fmt.Printf("修改后的颜色: %s\n", c.color)
	c.changeStrs1([]int{1, 2})
	fmt.Printf("changeStrs1修改后的数组: %v\n", c.strs)
	c.changeStrs2([]int{1, 2})
	fmt.Printf("changeStrs2修改后的数组: %v\n", c.strs)

	c.addStrs1()
	fmt.Printf("addStr1修改后的数组: %v\n", c.strs)
	c.addStrs2()
	fmt.Printf("addStr2修改后的数组: %v\n", c.strs)
}
