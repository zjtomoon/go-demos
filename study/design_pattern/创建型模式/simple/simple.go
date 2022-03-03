package simplefactory

import "fmt"

// 简单工厂模式

// go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
// NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi,%s", name)
}

type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello,%s", name)
}
