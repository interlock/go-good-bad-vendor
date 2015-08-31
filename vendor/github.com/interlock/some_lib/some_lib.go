package some_lib

import (
	"github.com/interlock/bad_lib"
	"github.com/interlock/good_lib"
)

// GetBadLib returns an instance of the bad lib.
// If no vendor is present, go will look back one vendor level until it does find it (or not).
// If vendor is present, that is the code it will use
func GetBadLib() *bad_lib.BadLib {
	return bad_lib.NewBadLib()
}

// GetGoodLib returns an instance of a good lib
// Same vendor behaviours apply
func GetGoodLib() good_lib.GoodLib {
	return good_lib.NewGoodLib()
}
