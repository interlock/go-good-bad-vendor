package main

import (
	"fmt"
	good "github.com/interlock/good_lib"
	some "github.com/interlock/some_lib"
)

func doGood(g good.GoodLib) {
	g.SetGood("love it")
	fmt.Printf("This was set: %s\n", g.Good())
}

func main() {
	g := good.NewGoodLib()
	doGood(g)
	og := some.GetGoodLib()
	doGood(og)
}
