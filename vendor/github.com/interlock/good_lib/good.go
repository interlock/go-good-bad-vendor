package good_lib

// GoodLib defines an interface to publicly expose
type GoodLib interface {
	SetGood(bad string)
	Good() string
}

// goodLib is an internal implementation of our interface
type goodLib struct {
	good string
}

// NewGoodLib returns our private implementation, but cast to the interface
func NewGoodLib() GoodLib {
	return &goodLib{}
}

// SetGood sets a value
func (g *goodLib) SetGood(good string) {
	g.good = good
}

// Good returns the value
func (g goodLib) Good() string {
	return g.good
}
