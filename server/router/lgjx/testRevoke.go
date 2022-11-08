package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRevokeRouter struct {
}

func (s *TestRevokeRouter) InitRevokeRouter(Router *gin.RouterGroup) {
	testRevokeRouter := Router.Group("testRevoke").Use(middleware.OperationRecord())
	testRevokeRouterWithoutRecord := Router.Group("testRevoke")
	var testRevokeApi = v1.ApiGroupApp.LgjxTestApiGroup.TestRevokeApi
	{
		testRevokeRouter.POST("createRevoke", testRevokeApi.CreateRevoke)             // 新建Revoke
		testRevokeRouter.DELETE("deleteRevoke", testRevokeApi.DeleteRevoke)           // 删除Revoke
		testRevokeRouter.DELETE("deleteRevokeByIds", testRevokeApi.DeleteRevokeByIds) // 批量删除Revoke
		testRevokeRouter.PUT("updateRevoke", testRevokeApi.UpdateRevoke)              // 更新Revoke
	}
	{
		testRevokeRouterWithoutRecord.GET("findRevoke", testRevokeApi.FindRevoke)       // 根据ID获取Revoke
		testRevokeRouterWithoutRecord.GET("getRevokeList", testRevokeApi.GetRevokeList) // 获取Revoke列表
	}
}
