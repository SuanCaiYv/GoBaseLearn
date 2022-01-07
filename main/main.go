package main

import (
	"fmt"
)

/*
之前写了一份base版放在另一台电脑里了，现在重新写一份简略版的基础语法
*/

type MyInterface interface {
	Print()
	Show()
}

type MyStruct struct {
}

func (m MyStruct) Print() {}

func (m MyStruct) Show() {}

func main() {
	var a interface{}
	a = MyStruct{}
	fmt.Println(a.(MyStruct))
}
