package main

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
)

func main() {
	fmt.Println("封装函数调用")
	id := "3"
	username := "demo01"
	err := Audit(id, username)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("直接调用")

	cmd := exec.Command("StartMysqlService.out", "5", "demo01")
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		fmt.Println("err :", err)
	}

	LS("-l")
}

func Audit(id string, username string) error {
	arg1 := id
	arg2 := username
	cmd := exec.Command("StartMysqlService.out", arg1, arg2)
	fmt.Println(reflect.TypeOf(id))
	fmt.Println(reflect.TypeOf(username))
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("err :", err)
	}
	return err
}

func LS(arg string) {
	cmd := exec.Command("ls", arg)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("err:", err)
	}
}
