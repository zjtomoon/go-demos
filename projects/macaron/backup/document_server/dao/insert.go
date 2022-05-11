package dao

import (
	"context"
	"fmt"

	"document_server/models"
	"gitee.com/chunanyong/zorm"
)

func InsertUser(username string, password string) {
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (i interface{}, e error) {
		user := models.NewUser(username, password)
		_, err := zorm.Insert(ctx, user)
		return nil, err
	})
	if err != nil {
		fmt.Println("failed to insert user,err:", err)
		return
	}
}
