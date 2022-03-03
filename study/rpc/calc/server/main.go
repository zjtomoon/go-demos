package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

type Arith struct {
}

type ArithRequest struct {
	A, B int
}

type ArithResponse struct {
	Pro int //乘积
	Quo int //商
	Rem int // 余数
}

// 乘法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	rect := new(Arith)
	// 注册rpc服务
	rpc.Register(rect)
	//服务处理绑定到http协议上
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
