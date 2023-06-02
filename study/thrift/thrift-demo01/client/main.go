package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"os"
	"thrift-demo01-client/gen-go/echo"
	"time"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport := thrift.NewTSocketConf("127.0.0.1:9898", &thrift.TConfiguration{
		ConnectTimeout: 10 * time.Second,
		SocketTimeout:  10 * time.Second,
	})

	useTransport, err := transportFactory.GetTransport(transport)
	client := echo.NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error openning socket to 127.0.0.1:9898", " ", err)
		os.Exit(1)
	}

	defer transport.Close()

	var ctx = context.Background()

	req := &echo.EchoReq{Msg: "You are welcome."}
	res, err := client.Echo(ctx, req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}
	log.Println("response:", res.Msg)
	fmt.Println("well done")

	num1 := &echo.Num{ID: 1}
	num2 := &echo.Num{ID: 2}
	num, err := client.Add(ctx, num1, num2)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}
	log.Println("result = ", num.ID)
	fmt.Println("well done")

}
