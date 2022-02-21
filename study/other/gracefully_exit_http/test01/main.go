package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Golang 如何优雅的终止一个服务
// golang http.Server结构体有一个终止服务的方法Shutdown，其go doc如下。
// 1、func (srv *Server) Shutdown(ctx context.Context) error

// 2、func Notify(c chan<- os.Signal, sig ...os.Signal)

//接下来我们使用如上signal.Notify结合http.Server的Shutdown方法实现服务优雅的终止。
// 如下代码，Handler与文章开始时的处理逻辑一样，其会在2s后返回hello。
// 创建一个http.Server实例，指定端口与Handler。
// 声明一个processed chan，其用来保证服务优雅的终止后再退出主goroutine。
// 新启一个goroutine，其会监听os.Interrupt信号，一旦服务被中断即调用服务的Shutdown方法，
// 确保活跃连接的正常返回（本代码使用的Context超时时间为3s，大于服务Handler的处理时间，所以不会超时）。
// 处理完成后，关闭processed通道，最后主goroutine退出。

var addr = flag.String("server addr", ":8080", "server address")

func main() {

	// handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "hello")
	})

	// server
	srv := http.Server{
		Addr:    *addr,
		Handler: handler,
	}

	// make sure idle connections returned
	processed := make(chan struct{})

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown failed,err : %v\n", err)
		}
		log.Fatalln("server gracefully shutdown")
		close(processed)
	}()

	//serve
	err := srv.ListenAndServe()
	if http.ErrServerClosed != err {
		log.Fatalf("server not gracefully shutdown,err : %v\n", err)
	}

	// waiting for goroutine above processed
	<-processed
}
