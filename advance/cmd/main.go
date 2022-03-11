package main

import (
	"fmt"
)

func main() {
	fmt.Println(test())
}

var str = "aaa"

func test() (v string) {
	defer func(p *string) {
		*p = str
	}(&v)
	return v
}
