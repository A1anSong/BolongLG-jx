package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestApplyRouter struct {
}

func (s *TestApplyRouter) InitApplyRouter(Router *gin.RouterGroup) {
	testApplyRouter := Router.Group("testApply").Use(middleware.OperationRecord())
	testApplyRouterWithoutRecord := Router.Group("testApply")
	var testApplyApi = v1.ApiGroupApp.LgjxTestApiGroup.TestApplyApi
	{
		testApplyRouter.POST("createApply", testApplyApi.CreateApply)             // 新建Apply
		testApplyRouter.DELETE("deleteApply", testApplyApi.DeleteApply)           // 删除Apply
		testApplyRouter.DELETE("deleteApplyByIds", testApplyApi.DeleteApplyByIds) // 批量删除Apply
		testApplyRouter.PUT("updateApply", testApplyApi.UpdateApply)              // 更新Apply
	}
	{
		testApplyRouterWithoutRecord.GET("findApply", testApplyApi.FindApply)       // 根据ID获取Apply
		testApplyRouterWithoutRecord.GET("getApplyList", testApplyApi.GetApplyList) // 获取Apply列表
	}
}
