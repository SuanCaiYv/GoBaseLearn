package main

import (
	"fmt"
	"unsafe"
)

/*
之前写了一份base版放在另一台电脑里了，现在重新写一份简略版的基础语法
*/

func main() {
	type T struct {
		a int8
		b int32
		c int16
		d int64
		e int16
	}
	t := T{}
	fmt.Println(unsafe.Alignof(t.a))
	fmt.Println(unsafe.Alignof(t.b))
	fmt.Println(unsafe.Alignof(t.c))
	fmt.Println(unsafe.Alignof(t.d))
	fmt.Println(unsafe.Sizeof(t))
}
