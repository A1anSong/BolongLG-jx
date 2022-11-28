package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestOrderRouter struct {
}

func (s *TestOrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	testOrderRouter := Router.Group("testOrder").Use(middleware.OperationRecord())
	testOrderRouterWithoutRecord := Router.Group("testOrder")
	var testOrderApi = v1.ApiGroupApp.LgjxTestApiGroup.TestOrderApi
	{
		testOrderRouter.POST("createOrder", testOrderApi.CreateOrder)             // 新建Order
		testOrderRouter.DELETE("deleteOrder", testOrderApi.DeleteOrder)           // 删除Order
		testOrderRouter.DELETE("deleteOrderByIds", testOrderApi.DeleteOrderByIds) // 批量删除Order
		testOrderRouter.PUT("updateOrder", testOrderApi.UpdateOrder)              // 更新Order
	}
	{
		testOrderRouterWithoutRecord.GET("findOrder", testOrderApi.FindOrder)                         // 根据ID获取Order
		testOrderRouterWithoutRecord.GET("getOrderList", testOrderApi.GetOrderList)                   // 获取Order列表
		testOrderRouterWithoutRecord.GET("getOrderStatisticData", testOrderApi.GetOrderStatisticData) // 获取Order统计数据
		testOrderRouterWithoutRecord.GET("exportExcel", testOrderApi.ExportExcel)                     // 导出Order数据到excel
	}
}
