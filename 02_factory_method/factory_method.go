package factoryMethod

type Operator interface {
	SetA(a float64)
	SetB(b float64)
	Result() float64
}

type OperatorFactory interface {
	Create() Operator
}

// 定义一个实现Operator接口的基类
type OperatorBase struct {
	a, b float64
}

func (ob *OperatorBase) SetA(a float64)  {
	ob.a = a
}

func (ob *OperatorBase) SetB(b float64)  {
	ob.b = b
}

// PlusOperator工厂类
type PlusOperatorFactory struct {}

// Operator 的加法实现
// 继承自OperatorBase基类
type PlusOperator struct {
	*OperatorBase
}

func (po *PlusOperator) Result() float64 {
	return po.a + po.b
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}
