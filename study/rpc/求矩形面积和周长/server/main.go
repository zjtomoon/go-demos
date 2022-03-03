package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

type Rect struct {
}

// 矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

// 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

func main() {
	rect := new(Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Panicln(err)
	}
}

// rpc程序的四个基本条件

// 结构体字段首字母要大写，可以别人调用
//
//函数名必须首字母大写
//
//函数第一参数是接收参数，第二个参数是返回给客户端的参数，必须是指针类型
//
//函数还必须有一个返回值error
