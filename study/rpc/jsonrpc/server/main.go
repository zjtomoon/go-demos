package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Params struct {
	Width, Height int
}

type Rect struct {
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// net/rpc/jsonrpc库通过json格式编解码，支持跨语言调用
func main() {
	rpc.Register(new(Rect))
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panicln(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
