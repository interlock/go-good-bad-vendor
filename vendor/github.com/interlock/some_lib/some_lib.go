package some_lib

import (
	"github.com/interlock/bad_lib"
	"github.com/interlock/good_lib"
)

func GetBadLib() *bad_lib.BadLib {
	return bad_lib.NewBadLib()
}

func GetGoodLib() good_lib.GoodLib {
	return good_lib.NewGoodLib()
}
