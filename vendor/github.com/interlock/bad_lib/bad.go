package bad_lib

type BadLib struct {
	bad string
}

func NewBadLib() *BadLib {
	return &BadLib{}
}

func (b *BadLib) SetBad(bad string) {
	b.bad = bad
}

func (b BadLib) Bad() string {
	return b.bad
}
