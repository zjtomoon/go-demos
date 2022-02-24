package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Golang 中实现一个 HTTP SERVER 异常的简单，利用标准库 net/http 的实现仅需数行代码即可，
// 但是一个生产环境可用的 HTTP SERVER 还必须考虑更多的问题，其中如何实现优雅关闭 HTTP SERVER 是一个必须要处理的问题。
// 这里所说的 优雅 即是指在 HTTP SERVER 里监听特定的信号，并在接收到信号时做出相应的处理。中文世界里也常用平滑 一词代替，
// 说的都是同一个意思。
const addr = ":9527"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "just another http server")
	})
	//使用默认路由创建http server
	srv := http.Server{
		Addr:    addr,
		Handler: http.DefaultServeMux,
	}
	var wg sync.WaitGroup
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-exit
		wg.Add(1)
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		err := srv.Shutdown(ctx)
		if err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()
	fmt.Println("listen at " + addr)
	err := srv.ListenAndServe()
	fmt.Println("waiting for the remaining connections to finish...")
	wg.Wait()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	fmt.Println("gracefully shutdown the http server...")
}
