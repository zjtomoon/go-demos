package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// 查询命令所在目录，可用于判断操作系统是否能用该命令
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal("install fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)

	// cmd中执行的命令:ffmpeg -i psg.flv test.mp4

	cmd := exec.Command("ffmpeg", "-i", "psg.flv", "tes.mp4")

	err1 := cmd.Run()
	if err1 != nil {
		fmt.Println("err")
	}
}
