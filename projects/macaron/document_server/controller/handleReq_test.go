package controller

import (
	"fmt"
	"testing"
)

func TestHandleReq(t *testing.T) {
	req := "login:szf+123"
	reqtype, username, password := HandleReq(req)
	fmt.Println("reqtype:", reqtype)
	fmt.Println("username:", username)
	fmt.Println("password:", password)
}
