package main

import (
	"fmt"
	"reflect"
)

//基本类型反射
func basicReflect() {
	x := 1
	type1 := reflect.TypeOf(x) //获取x的类型
	fmt.Println("type1: ", type1)

	xv := reflect.ValueOf(x)
	fmt.Println("value: ", xv)
	type2 := xv.Type()
	fmt.Println("xv type: ", type2)
	kid := xv.Kind()
	fmt.Println("xv kind: ", kid)
	inter := xv.Interface()
	fmt.Println("xv Interface: ", inter)
	cs := xv.CanSet()
	fmt.Println("xv CanSet: ", cs)
	ele := xv.Elem()
	fmt.Println("xv Elem: ", ele)
	fmt.Println("xv Elem type: ", ele.Type())

	xv2 := reflect.ValueOf(&x)
	cs2 := xv2.CanSet()
	fmt.Println("xv2 CanSet: ", cs2)
	ele2 := xv2.Elem()
	fmt.Println("ele CanSet: ", ele2.CanSet())
	ele2.SetInt(22)
	fmt.Println("x 经过反射修改后的值：", x)

}

//定义一个结构体
type testStruct struct {
	name string
	age  int
}

//给结构体添加一个方法
func (t testStruct) Show() {
	fmt.Printf("name: %s, age: %d\n", t.name, t.age)
}

func (t *testStruct) SetAge(age int) {
	t.age = age
}

//结构体反射
func structReflect() {
	var ts *testStruct = &testStruct{"测试name", 20}
	r := reflect.ValueOf(ts) //获取结构体的反射对象
	fmt.Println("反射对象类型: ", r.Type())
	ele := r.Elem() //反射对象的内容(这里指的是反射对象指向的对象)
	fmt.Println("反射对象的内容：", ele)
	fmt.Println("反射对象的内容的类型：", ele.Type())

	//调用无法参数构造方法
	var args1 = make([]reflect.Value, 0) //长度为0的参数数组
	//var args1 = []reflect.Value{}      //长度为0的参数数组
	r.MethodByName("Show").Call(args1) //反射调用方法

	//调用有参数构造函数
	var s2 int
	s2 = 222
	var args2 = []reflect.Value{reflect.ValueOf(s2)}
	r.MethodByName("SetAge").Call(args2)
	fmt.Println("通过反射对象获取字段age的值: ", reflect.Indirect(r).FieldByName("age").Kind())
	fmt.Println("通过原始对象获取字段age的值: ", ts.age)

}

func main() {
	//basicReflect()
	structReflect()
}
