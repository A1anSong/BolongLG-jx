package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestDelayRouter struct {
}

func (s *TestDelayRouter) InitDelayRouter(Router *gin.RouterGroup) {
	testDelayRouter := Router.Group("testDelay").Use(middleware.OperationRecord())
	testDelayRouterWithoutRecord := Router.Group("testDelay")
	var testDelayApi = v1.ApiGroupApp.LgjxTestApiGroup.TestDelayApi
	{
		testDelayRouter.POST("createDelay", testDelayApi.CreateDelay)             // 新建Delay
		testDelayRouter.DELETE("deleteDelay", testDelayApi.DeleteDelay)           // 删除Delay
		testDelayRouter.DELETE("deleteDelayByIds", testDelayApi.DeleteDelayByIds) // 批量删除Delay
		testDelayRouter.PUT("updateDelay", testDelayApi.UpdateDelay)              // 更新Delay
	}
	{
		testDelayRouterWithoutRecord.GET("findDelay", testDelayApi.FindDelay)       // 根据ID获取Delay
		testDelayRouterWithoutRecord.GET("getDelayList", testDelayApi.GetDelayList) // 获取Delay列表
	}
}
