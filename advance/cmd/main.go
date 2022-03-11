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
