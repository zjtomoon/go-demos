package main

import (
	"document_server/dao"
	"gopkg.in/macaron.v1"
)

func main() {
	dao.InsertUser("user1", "123456")
	m := macaron.Classic()
	m.Get("/", func() string {
		return "hello"
	})
	m.Run(8080)
}
