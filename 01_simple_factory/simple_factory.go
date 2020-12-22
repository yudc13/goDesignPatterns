package simpleFactory

type API interface {
	Fetch(url string) string
}

type NetApi struct{}

// 实现API接口
func (netApi *NetApi) Fetch(url string) string {
	return "hi netApi " + url
}

type MockApi struct{}

func (mockApi *MockApi) Fetch(url string) string {
	return "hi mockApi " + url
}

// 工厂方法
func NewApi(t string) API {
	switch t {
	case "net":
		return &NetApi{}
	case "mock":
		return &MockApi{}
	default:
		return nil
	}
}
