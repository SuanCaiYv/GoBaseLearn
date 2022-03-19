package main

import (
	"fmt"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp", ":8080")
	accept, err := listen.Accept()
	if err != nil {
		return
	}
	fmt.Println(accept)
}
