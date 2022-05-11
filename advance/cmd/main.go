package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func more[T ~[]E, E constraints.Integer](slice T, factor E) []E { // 仅仅改了返回值
	tmp := make([]E, len(slice))
	for i, e := range slice {
		tmp[i] = e * factor
	}
	return tmp
}

type myInts []int

func (m myInts) print() {
	for _, e := range m {
		fmt.Printf("%d ", e)
	}
}

func main() {
	slice := myInts{1, 2, 3}
	tmp := myInts(more(slice, 2))
	tmp.print() // OK
}
