package main

import (
	"fmt"
	"io/ioutil"

	"github.com/json-iterator/go"
)

func main() {
	fmt.Println(Sysconfig)
	fmt.Println(Sysconfig.DBName)
}

var Sysconfig = &sysconfig{}

func init() {
	// 指定对应的json配置文件
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Sys config read err")
	}
	err = jsoniter.Unmarshal(b, Sysconfig)
	if err != nil {
		panic(err)
	}
}

type sysconfig struct {
	Port       string `json:"Port"`
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBIp       string `json:"DBIp"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
}