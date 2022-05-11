package controller

import (
	"fmt"
	"net"
	"os"
)

//todo : 发送json文件

// 参考: https://blog.csdn.net/weixin_42094659/article/details/101674891
// 参考: https://blog.csdn.net/cui_yonghua/article/details/95213747
// socket协议发送和接收：https://blog.csdn.net/weixin_45477086/article/details/122933131

func SendFile(path string, conn net.Conn) {
	// 以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer f.Close()

	// 读文件内容，读多少就发送多少内容
	buf := make([]byte, 1024*4)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err != nil {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err:", err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}
