package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestInvoiceRouter struct {
}

func (s *TestInvoiceRouter) InitInvoiceRouter(Router *gin.RouterGroup) {
	testInvoiceRouter := Router.Group("testInvoice").Use(middleware.OperationRecord())
	testInvoiceRouterWithoutRecord := Router.Group("testInvoice")
	var testInvoiceApi = v1.ApiGroupApp.LgjxTestApiGroup.TestInvoiceApi
	{
		testInvoiceRouter.POST("createInvoice", testInvoiceApi.CreateInvoice)             // 新建Invoice
		testInvoiceRouter.DELETE("deleteInvoice", testInvoiceApi.DeleteInvoice)           // 删除Invoice
		testInvoiceRouter.DELETE("deleteInvoiceByIds", testInvoiceApi.DeleteInvoiceByIds) // 批量删除Invoice
		testInvoiceRouter.PUT("updateInvoice", testInvoiceApi.UpdateInvoice)              // 更新Invoice
	}
	{
		testInvoiceRouterWithoutRecord.GET("findInvoice", testInvoiceApi.FindInvoice)       // 根据ID获取Invoice
		testInvoiceRouterWithoutRecord.GET("getInvoiceList", testInvoiceApi.GetInvoiceList) // 获取Invoice列表
	}
}
