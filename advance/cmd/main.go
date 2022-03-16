package main

import (
	"fmt"
)

type T struct {
	a int64
	b int64
	c int64
	d int64
}

func main() {
	PtrSize := 4 << (^uintptr(0) >> 63)
	uintptrMask := 1<<(8*PtrSize) - 1
	fmt.Println(uintptrMask & -1314)
	f()
}

func test1() (t T) {
	t = T{1, 2, 3, 4}
	return t
}

func test2() T {
	v := T{1, 2, 3, 4}
	return v
}

func f() {
	defer func() {
		fmt.Println("bbb")
		fmt.Println(recover())
	}()
	defer func() {
		fmt.Println("ccc")
	}()
	panic("aaa")
	fmt.Println("ddd")
}
