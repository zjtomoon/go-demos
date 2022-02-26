package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

func main() {
	fmt.Println("Connecting to mongodb using mgo")
	session := InitMongoSession()
	myDB := session.DB("temp")
	collects, err := myDB.CollectionNames()
	if err != nil {
		fmt.Println("CollectionNames-error:", err)
		os.Exit(0)
	}
	fmt.Println(collects)
}

func InitMongoSession() *mgo.Session {
	mHost := "127.0.0.1"
	mPort := "27017"
	//mDBName := "config"  //你要连接的表，也可以后面再选，都行
	mUsername := "admin"
	mPassword := "123456"

	session, err := mgo.Dial(mHost + ":" + mPort)

	if err != nil {
		fmt.Println("mgo.Dial-error:", err)
		os.Exit(0)
	}
	session.SetMode(mgo.Eventual, true)
	myDB := session.DB("admin") // 这里的关键是连接mongodb后，选择admin数据库，然后登录，确保账号密码无误之后，该连接就一直能用了
	err = myDB.Login(mUsername, mPassword)
	if err != nil {
		fmt.Println("Login-error:", err) //mgo在mongodb5.0以后无法连接

		/*
			MongoDB 5 支持SCRAM-SHA-1和SCRAM-SHA-256验证，mgo.v2默认是用SCRAM-SHA-1验证的，
			换个authMechanism=SCRAM-SHA-256验证后，就报出个另一个错误SASL support not enabled during build (-tags sasl)。

		*/
		os.Exit(0)
	}
	//myDB = session.DB(mDBName) //如果要在这里就选择数据库，这个myDB可以定义为全局变量
	session.SetPoolLimit(10)
	return session
}
