package simpleFactory

type Gretting interface {
	Say(content string) string
}

type GoodMorning struct{}

func (gm GoodMorning) Say(content string) string {
	return "hi good morning " + content
}

type GoodEvening struct{}

func (ge GoodEvening) Say(content string) string {
	return "hi good evening " + content
}

func NewGretting(t string) Gretting {
	switch t {
	case "morning":
		return GoodMorning{}
	case "evening":
		return GoodEvening{}
	default:
		return nil
	}
}