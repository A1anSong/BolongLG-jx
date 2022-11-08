package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestClaimRouter struct {
}

func (s *TestClaimRouter) InitClaimRouter(Router *gin.RouterGroup) {
	testClaimRouter := Router.Group("testClaim").Use(middleware.OperationRecord())
	testClaimRouterWithoutRecord := Router.Group("testClaim")
	var testClaimApi = v1.ApiGroupApp.LgjxTestApiGroup.TestClaimApi
	{
		testClaimRouter.POST("createClaim", testClaimApi.CreateClaim)             // 新建Claim
		testClaimRouter.DELETE("deleteClaim", testClaimApi.DeleteClaim)           // 删除Claim
		testClaimRouter.DELETE("deleteClaimByIds", testClaimApi.DeleteClaimByIds) // 批量删除Claim
		testClaimRouter.PUT("updateClaim", testClaimApi.UpdateClaim)              // 更新Claim
	}
	{
		testClaimRouterWithoutRecord.GET("findClaim", testClaimApi.FindClaim)       // 根据ID获取Claim
		testClaimRouterWithoutRecord.GET("getClaimList", testClaimApi.GetClaimList) // 获取Claim列表
	}
}
