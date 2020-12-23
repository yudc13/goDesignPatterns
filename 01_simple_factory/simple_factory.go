package simpleFactory

type Gretting interface {
	Say(content string) string
}

type goodMorning struct{}

func (gm goodMorning) Say(content string) string {
	return "hi good morning " + content
}

type goodEvening struct{}

func (ge goodEvening) Say(content string) string {
	return "hi good evening " + content
}

func NewGretting(t string) Gretting {
	switch t {
	case "morning":
		return goodMorning{}
	case "evening":
		return goodEvening{}
	default:
		return nil
	}
}