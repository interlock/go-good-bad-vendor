package bad_lib

// BadLib is a simple publicly exposed library
type BadLib struct {
	bad string
}

// NewBadLib creates a new instance of our struct
func NewBadLib() *BadLib {
	return &BadLib{}
}

// SetBad sets the bad value, which is really just fine... but bad by association
func (b *BadLib) SetBad(bad string) {
	b.bad = bad
}

// Bad returns the value set
func (b BadLib) Bad() string {
	return b.bad
}
