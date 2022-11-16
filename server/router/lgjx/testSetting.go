package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestSettingRouter struct {
}

func (s *TestSettingRouter) InitSettingRouter(Router *gin.RouterGroup) {
	testSettingRouter := Router.Group("testSetting").Use(middleware.OperationRecord())
	testSettingRouterWithoutRecord := Router.Group("testSetting")
	var testSettingApi = v1.ApiGroupApp.LgjxTestApiGroup.TestSettingApi
	{
		testSettingRouter.POST("createSetting", testSettingApi.CreateSetting)             // 新建Setting
		testSettingRouter.DELETE("deleteSetting", testSettingApi.DeleteSetting)           // 删除Setting
		testSettingRouter.DELETE("deleteSettingByIds", testSettingApi.DeleteSettingByIds) // 批量删除Setting
		testSettingRouter.PUT("updateSetting", testSettingApi.UpdateSetting)              // 更新Setting
	}
	{
		testSettingRouterWithoutRecord.GET("findSetting", testSettingApi.FindSetting)       // 根据ID获取Setting
		testSettingRouterWithoutRecord.GET("getSettingList", testSettingApi.GetSettingList) // 获取Setting列表
	}
}
