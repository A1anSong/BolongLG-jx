package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestLogoutRouter struct {
}

func (s *TestLogoutRouter) InitLogoutRouter(Router *gin.RouterGroup) {
	testLogoutRouter := Router.Group("testLogout").Use(middleware.OperationRecord())
	testLogoutRouterWithoutRecord := Router.Group("testLogout")
	var testLogoutApi = v1.ApiGroupApp.LgjxTestApiGroup.TestLogoutApi
	{
		testLogoutRouter.POST("createLogout", testLogoutApi.CreateLogout)             // 新建Logout
		testLogoutRouter.DELETE("deleteLogout", testLogoutApi.DeleteLogout)           // 删除Logout
		testLogoutRouter.DELETE("deleteLogoutByIds", testLogoutApi.DeleteLogoutByIds) // 批量删除Logout
		testLogoutRouter.PUT("updateLogout", testLogoutApi.UpdateLogout)              // 更新Logout
	}
	{
		testLogoutRouterWithoutRecord.GET("findLogout", testLogoutApi.FindLogout)       // 根据ID获取Logout
		testLogoutRouterWithoutRecord.GET("getLogoutList", testLogoutApi.GetLogoutList) // 获取Logout列表
	}
}
