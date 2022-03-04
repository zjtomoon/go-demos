package factorymethod

// 工厂方法模式
// 工厂方法模式使用子类的方式延迟生成对象到子类中实现。

// Operator 是被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

// SetA 设置A

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory 是PlusOperator的工厂类
type PlusOperatorFactory struct{}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

type MinusOperator struct {
	*OperatorBase
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

func (o MinusOperator) Result() int {
	return o.a - o.b
}
