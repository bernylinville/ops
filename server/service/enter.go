package service

import (
	"github.com/bernylinville/ops/server/service/example"
	"github.com/bernylinville/ops/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
