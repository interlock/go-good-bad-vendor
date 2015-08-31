package good_lib

type GoodLib interface {
	SetGood(bad string)
	Good() string
}

type goodLib struct {
	good string
}

func NewGoodLib() GoodLib {
	return &goodLib{}
}

func (g *goodLib) SetGood(good string) {
	g.good = good
}

func (g goodLib) Good() string {
	return g.good
}
