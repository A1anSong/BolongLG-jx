package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/lgjx"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	JRAPI     lgjx.JRAPIRouter
	JRAPITest lgjx.TESTJRAPIRouter
	Lgjx      lgjx.RouterGroup
	LgjxTest  lgjx.TestRouterGroup
}

var RouterGroupApp = new(RouterGroup)
