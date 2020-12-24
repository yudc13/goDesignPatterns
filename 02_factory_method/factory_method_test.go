package factoryMethod

import "testing"

func TestPlusOperator(t *testing.T) {
	factory := PlusOperatorFactory{}
	operator := factory.Create()
	operator.SetA(3.5)
	operator.SetA(4.5)
	t.Log(operator.Result())
}

func TestWulingCar(t *testing.T)  {
	wuling := WulingFactory{}.Create()
	name := wuling.GetName()
	if name != "五菱宏光" {
		t.Errorf("test WulingFactory error")
	}
}

func TestBMWCar(t *testing.T)  {
	bmw := BMWFactory{}.Create()
	name := bmw.GetName()
	if name != "宝马" {
		t.Errorf("test BMWFactory error")
	}
}
