package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestPayRouter struct {
}

func (s *TestPayRouter) InitPayRouter(Router *gin.RouterGroup) {
	testPayRouter := Router.Group("testPay").Use(middleware.OperationRecord())
	testPayRouterWithoutRecord := Router.Group("testPay")
	var testPayApi = v1.ApiGroupApp.LgjxTestApiGroup.TestPayApi
	{
		testPayRouter.POST("createPay", testPayApi.CreatePay)             // 新建Pay
		testPayRouter.DELETE("deletePay", testPayApi.DeletePay)           // 删除Pay
		testPayRouter.DELETE("deletePayByIds", testPayApi.DeletePayByIds) // 批量删除Pay
		testPayRouter.PUT("updatePay", testPayApi.UpdatePay)              // 更新Pay
	}
	{
		testPayRouterWithoutRecord.GET("findPay", testPayApi.FindPay)       // 根据ID获取Pay
		testPayRouterWithoutRecord.GET("getPayList", testPayApi.GetPayList) // 获取Pay列表
	}
}
