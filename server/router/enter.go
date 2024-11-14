package router

import (
	"github.com/bernylinville/ops/server/router/example"
	"github.com/bernylinville/ops/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
