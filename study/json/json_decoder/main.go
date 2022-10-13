package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	AppName string `json:"appname"`
	Port    int    `json:"port"`
}

func main() {
	file,_ := os.Open("./config.json")
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:",err)
	}
	fmt.Println(conf.Port)
}
