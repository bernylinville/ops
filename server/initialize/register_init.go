package initialize

import (
	_ "github.com/bernylinville/ops/server/source/example"
	_ "github.com/bernylinville/ops/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
