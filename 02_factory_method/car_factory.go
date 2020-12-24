package factoryMethod

type CarFactory interface {
	Create() Car
}

type WulingFactory struct {}

type BMWFactory struct {}

func (WulingFactory) Create() Car {
	return Wuling{&CarBase{
		Name: "五菱宏光",
	}}
}

func (BMWFactory) Create() Car {
	return BMW{&CarBase{
		Name: "宝马",
	}}
}

