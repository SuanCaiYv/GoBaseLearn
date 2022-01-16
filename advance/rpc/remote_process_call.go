package rpc

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

type Server1 struct{}

func (s *Server1) AppendStr(request string, reply *string) error {
	*reply = "server reply: " + request
	return nil
}

func RunRpc1() {
	_ = rpc.RegisterName("Service1", &Server1{})
	go func() {
		serverSocket, _ := net.Listen("tcp", ":8191")
		for {
			socket, _ := serverSocket.Accept()
			rpc.ServeConn(socket)
		}
	}()
	time.Sleep(time.Duration(100) * time.Millisecond)
	go func() {
		client, _ := rpc.Dial("tcp", ":8191")
		var reply string
		_ = client.Call("Service1.AppendStr", "aaa", &reply)
		fmt.Println(reply)
	}()
	<-make(chan struct{})
}
