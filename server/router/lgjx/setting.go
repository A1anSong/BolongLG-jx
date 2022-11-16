package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SettingRouter struct {
}

// InitSettingRouter 初始化 Setting 路由信息
func (s *SettingRouter) InitSettingRouter(Router *gin.RouterGroup) {
	settingRouter := Router.Group("setting").Use(middleware.OperationRecord())
	settingRouterWithoutRecord := Router.Group("setting")
	var settingApi = v1.ApiGroupApp.LgjxApiGroup.SettingApi
	{
		settingRouter.POST("createSetting", settingApi.CreateSetting)             // 新建Setting
		settingRouter.DELETE("deleteSetting", settingApi.DeleteSetting)           // 删除Setting
		settingRouter.DELETE("deleteSettingByIds", settingApi.DeleteSettingByIds) // 批量删除Setting
		settingRouter.PUT("updateSetting", settingApi.UpdateSetting)              // 更新Setting
	}
	{
		settingRouterWithoutRecord.GET("findSetting", settingApi.FindSetting)       // 根据ID获取Setting
		settingRouterWithoutRecord.GET("getSettingList", settingApi.GetSettingList) // 获取Setting列表
	}
}
