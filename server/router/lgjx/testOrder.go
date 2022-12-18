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
		testOrderRouter.PUT("approveApply", testOrderApi.ApproveApply)            // 审批同意申请
		testOrderRouter.PUT("rejectApply", testOrderApi.RejectApply)              // 审批拒绝申请
		testOrderRouter.PUT("approveDelay", testOrderApi.ApproveDelay)            // 审批同意延期
		testOrderRouter.PUT("rejectDelay", testOrderApi.RejectDelay)              // 审批拒绝延期
		testOrderRouter.PUT("approveRefund", testOrderApi.ApproveRefund)          // 审批同意退函
		testOrderRouter.PUT("rejectRefund", testOrderApi.RejectRefund)            // 审批拒绝退函
		testOrderRouter.PUT("approveClaim", testOrderApi.ApproveClaim)            // 审批同意理赔
		testOrderRouter.PUT("rejectClaim", testOrderApi.RejectClaim)              // 审批拒绝理赔
		testOrderRouter.PUT("openLetter", testOrderApi.OpenLetter)                // 提交开函申请
	}
	{
		testOrderRouterWithoutRecord.GET("findOrder", testOrderApi.FindOrder)                         // 根据ID获取Order
		testOrderRouterWithoutRecord.GET("getOrderList", testOrderApi.GetOrderList)                   // 获取Order列表
		testOrderRouterWithoutRecord.GET("getOrderStatisticData", testOrderApi.GetOrderStatisticData) // 获取Order统计数据
		testOrderRouterWithoutRecord.GET("exportExcel", testOrderApi.ExportExcel)                     // 导出Order数据到excel
	}
}
