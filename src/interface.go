package src

import (
	"fmt"
)

type display interface {
	display0()
}

type show interface {
	show0()
}

type tmpA struct {
}

type tmpB struct {
}

// 关于实现方法，我说一些不太容易注意到的细节。
// 如果某个类型实现了接口A，但是任一方法的接收者类型是指针类型的话，则接口A定义的变量只能是一个指针
// 如果某个实现了接口A的类型的所有方法的接收者类型都不是指针类型的话，则接口A定一个变量可以是指针，也可以是普通值
// 所以类型为interface但是指向类型的变量，至少是一个指针类型，interface具有同时是指针和普通变量两种属性
func (t *tmpA) display0() {
	fmt.Println("display tmpA")
}

func (t *tmpA) show0() {
	fmt.Println("show tmpA")
}

func (t *tmpB) show0() {
	fmt.Println("show tmpB")
}

// x.(T)语法指出T是否实现了x接口，如果实现了，返回实际类型下的变量，即父类强制转换到子类
// 如果没有实现，则标志位为false，返回nil
func fn71() {
	var v1 display
	// v1的静态类型是display，动态类型是*tmpA，值是tmpA的字段
	v1 = &tmpA{}
	v2, ok := v1.(*tmpA)
	if !ok {
		fmt.Println("not tmpB")
	} else {
		// v2是v1的动态类型下的变量，所以拥有tmpA的全部方法
		v1.display0()
		v2.show0()
	}
	var v3 show
	v3 = &tmpB{}
	_, ok = v3.(display)
	if !ok {
		fmt.Println("not match")
	}
	_, ok = v3.(*tmpA)
	if !ok {
		fmt.Println("not tmpA")
	}
}

func fn72(x show) {
	// 当()内为type关键字是，会变成对实际类型的解析，对比Java中的instanceof关键字
	switch x.(type) {
	case *tmpA:
		fmt.Println("is tmpA")
	case *tmpB:
		fmt.Println("is tmpB")
	}
}

func fn73() {
	v1 := &tmpA{}
	v2 := &tmpB{}
	fn72(v1)
	fn72(v2)
}

// ShowInterface Java中拥有多态，也拥有静态类型和动态类型，现在来对比理解一下
// 接口是一个方法的集合，两个接口的类型相不相同取决于它们所有的方法是否相同
// 得益于多态，一个静态类型为接口类型的变量还有一个动态类型为某个struct的实际类型
// 通过x.(T)语法可以获取到接口类型变量的实际类型，并触发强制类型转换到实际类型(如果匹配，然后以返回值的形式返回转换结果)
func ShowInterface() {
	fn71()
	fn73()
}
