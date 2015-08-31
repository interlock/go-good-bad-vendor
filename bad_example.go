package main

import (
	"fmt"
	bad "github.com/interlock/bad_lib"
	some "github.com/interlock/some_lib"
)

func doBad(g *bad.BadLib) {
	g.SetBad("love it")
	fmt.Printf("This was set: %s\n", g.Bad())
}

func main() {
	g := bad.NewBadLib()
	doBad(g)
	og := some.GetBadLib()
	doBad(og)
}
