package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestTemplateRouter struct {
}

func (s *TestTemplateRouter) InitTemplateRouter(Router *gin.RouterGroup) {
	testTemplateRouter := Router.Group("testTemplate").Use(middleware.OperationRecord())
	testTemplateRouterWithoutRecord := Router.Group("testTemplate")
	var testTemplateApi = v1.ApiGroupApp.LgjxTestApiGroup.TestTemplateApi
	{
		testTemplateRouter.POST("createTemplate", testTemplateApi.CreateTemplate)             // 新建Template
		testTemplateRouter.DELETE("deleteTemplate", testTemplateApi.DeleteTemplate)           // 删除Template
		testTemplateRouter.DELETE("deleteTemplateByIds", testTemplateApi.DeleteTemplateByIds) // 批量删除Template
		testTemplateRouter.PUT("updateTemplate", testTemplateApi.UpdateTemplate)              // 更新Template
	}
	{
		testTemplateRouterWithoutRecord.GET("findTemplate", testTemplateApi.FindTemplate)       // 根据ID获取Template
		testTemplateRouterWithoutRecord.GET("getTemplateList", testTemplateApi.GetTemplateList) // 获取Template列表
	}
}
