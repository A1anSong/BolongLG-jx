package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/lgjx"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	LgjxServiceGroup     lgjx.ServiceGroup
	LgjxTestServiceGroup lgjx.TestServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
