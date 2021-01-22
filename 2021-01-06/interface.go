package main

import "fmt"

//定义一个接口
type cat interface {
	miao(str string)
	changeName(newName string)
}

type whiteCat struct {
	name string
	age  int
}

func (c *whiteCat) miao(str string) {
	fmt.Printf("whiteCat miao: %s\n", str)
}

func (c *whiteCat) changeName(n string) {
	c.name = n
}

type integer int

func (a integer) Less(b integer) bool {
	return a < b
}
func (a *integer) Add(b string) {

}

type lessAdder interface {
	Less(b integer) bool
	Add(b string)
}

func main() {
	var wc *whiteCat = &whiteCat{"白猫", 2}
	wc.miao("123")
	var cc cat = wc
	cc.miao("456")
	//接口查询
	wccc, ok := cc.(*whiteCat) //判断接口对象cc是否为whiteCat的一个实例
	fmt.Println(ok)
	fmt.Printf("%T\n", wccc)
	fmt.Println(wccc == wc)

	//接口类型查询
	switch cc.(type) { //询问接口指向的对象实例的类型，必须放在switch中
	case *whiteCat:
		fmt.Println("是whiteCat类型")
	default:
	}

	var a integer = 1
	var b lessAdder = &a
	fmt.Println(b)

	var ccc cat = new(whiteCat)
	ccc.miao("123")

}
