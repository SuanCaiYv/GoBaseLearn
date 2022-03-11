package main

import "fmt"

type T struct {
	a int64
	b int64
	c int64
	d int64
}

func main() {
	v1 := test1()
	v2 := test2()
	fmt.Println(v1, v2)
}

func test1() (t T) {
	t = T{1, 2, 3, 4}
	return t
}

func test2() T {
	v := T{1, 2, 3, 4}
	return v
}
