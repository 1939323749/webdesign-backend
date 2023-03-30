package grpc

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (p *HelloService) Hello(requst string, reply *string) error {
	*reply = "hello " + requst
	return nil
}
func HelloInit() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
	}
	rpc.ServeConn(conn)
}
