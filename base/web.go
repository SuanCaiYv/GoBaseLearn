package base

import "net"

func simple() {
	listener, _ := net.Listen("ipv4", ":8190")
	connection, _ := listener.Accept()
	connection.Close()
}
