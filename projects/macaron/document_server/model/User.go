package model

// todo: 使用雪花算法生成id
type User struct {
	Id        int64  `gorm:"primaryKey column:Id"`
	UserName  string `gorm:"UserName"`
	Password  string `gorm:"Password"`
	FileTable string `FileTable:"FileTable"`
}

// todo:添加filetable
func NewUser(id int64, username string, password string) *User {
	return &User{
		Id:       id,
		UserName: username,
		Password: password,
	}
}
