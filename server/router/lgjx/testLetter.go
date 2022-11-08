package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestLetterRouter struct {
}

func (s *TestLetterRouter) InitLetterRouter(Router *gin.RouterGroup) {
	testLetterRouter := Router.Group("testLetter").Use(middleware.OperationRecord())
	testLetterRouterWithoutRecord := Router.Group("testLetter")
	var testLetterApi = v1.ApiGroupApp.LgjxTestApiGroup.TestLetterApi
	{
		testLetterRouter.POST("createLetter", testLetterApi.CreateLetter)             // 新建Letter
		testLetterRouter.DELETE("deleteLetter", testLetterApi.DeleteLetter)           // 删除Letter
		testLetterRouter.DELETE("deleteLetterByIds", testLetterApi.DeleteLetterByIds) // 批量删除Letter
		testLetterRouter.PUT("updateLetter", testLetterApi.UpdateLetter)              // 更新Letter
	}
	{
		testLetterRouterWithoutRecord.GET("findLetter", testLetterApi.FindLetter)       // 根据ID获取Letter
		testLetterRouterWithoutRecord.GET("getLetterList", testLetterApi.GetLetterList) // 获取Letter列表
	}
}
