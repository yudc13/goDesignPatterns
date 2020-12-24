package factoryMethod

type Car interface {
	GetName() string
}

type CarBase struct {
	Name string
}

type Wuling struct {
	*CarBase
}

type BMW struct {
	*CarBase
}

func (wl Wuling) GetName() string {
	return wl.Name
}

func (bmw BMW) GetName() string {
	return bmw.Name
}