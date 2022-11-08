package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestRefundRouter struct {
}

func (s *TestRefundRouter) InitRefundRouter(Router *gin.RouterGroup) {
	testRefundRouter := Router.Group("testRefund").Use(middleware.OperationRecord())
	testRefundRouterWithoutRecord := Router.Group("testRefund")
	var testRefundApi = v1.ApiGroupApp.LgjxTestApiGroup.TestRefundApi
	{
		testRefundRouter.POST("createRefund", testRefundApi.CreateRefund)             // 新建Refund
		testRefundRouter.DELETE("deleteRefund", testRefundApi.DeleteRefund)           // 删除Refund
		testRefundRouter.DELETE("deleteRefundByIds", testRefundApi.DeleteRefundByIds) // 批量删除Refund
		testRefundRouter.PUT("updateRefund", testRefundApi.UpdateRefund)              // 更新Refund
	}
	{
		testRefundRouterWithoutRecord.GET("findRefund", testRefundApi.FindRefund)       // 根据ID获取Refund
		testRefundRouterWithoutRecord.GET("getRefundList", testRefundApi.GetRefundList) // 获取Refund列表
	}
}
