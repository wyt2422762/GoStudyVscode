package main

import "fmt"

type base struct{
	name string
	age uint8
}

func (b *base) show(){
	fmt.Printf("name = %s, age= %d\n", b.name, b.age)
}

type child1 struct{
	base
}

type child2 struct{
	*base
}

func main(){
	var c1 child1 //定义一个子类

	fmt.Println(c1)

	c1.age = 1
	c1.name = "子类"

	c1.base.show()
	c1.show() //该方法从父类继承，同 c.base.show()

	fmt.Println("-----------------------------")

	var bs base
	bs.age = 1
	bs.name = "父类"

	var c2 child2
	fmt.Println(c2)
	//c2.age = 2
	//c2.name = "子类2"
	c2.base = &bs
	
	c2.base.show()
	c2.show()

}