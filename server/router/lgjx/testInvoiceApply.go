package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestInvoiceApplyRouter struct {
}

func (s *TestInvoiceApplyRouter) InitInvoiceApplyRouter(Router *gin.RouterGroup) {
	testInvoiceApplyRouter := Router.Group("testInvoiceApply").Use(middleware.OperationRecord())
	testInvoiceApplyRouterWithoutRecord := Router.Group("testInvoiceApply")
	var testInvoiceApplyApi = v1.ApiGroupApp.LgjxTestApiGroup.TestInvoiceApplyApi
	{
		testInvoiceApplyRouter.POST("createInvoiceApply", testInvoiceApplyApi.CreateInvoiceApply)             // 新建InvoiceApply
		testInvoiceApplyRouter.DELETE("deleteInvoiceApply", testInvoiceApplyApi.DeleteInvoiceApply)           // 删除InvoiceApply
		testInvoiceApplyRouter.DELETE("deleteInvoiceApplyByIds", testInvoiceApplyApi.DeleteInvoiceApplyByIds) // 批量删除InvoiceApply
		testInvoiceApplyRouter.PUT("updateInvoiceApply", testInvoiceApplyApi.UpdateInvoiceApply)              // 更新InvoiceApply
	}
	{
		testInvoiceApplyRouterWithoutRecord.GET("findInvoiceApply", testInvoiceApplyApi.FindInvoiceApply)       // 根据ID获取InvoiceApply
		testInvoiceApplyRouterWithoutRecord.GET("getInvoiceApplyList", testInvoiceApplyApi.GetInvoiceApplyList) // 获取InvoiceApply列表
	}
}
