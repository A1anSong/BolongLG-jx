package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TESTJRAPIRouter struct {
}

func (s *TESTJRAPIRouter) InitJRAPIRouter(Router *gin.RouterGroup) {
	testJRAPIRouter := Router.Group("test/lg").Use(middleware.JRValidate())
	testJRAPIRouterPublic := Router.Group("test/lg")
	testJRApi := v1.ApiGroupApp.LgjxTestApiGroup.TestJRAPI
	{
		testJRAPIRouter.POST("apply", testJRApi.Apply)
		testJRAPIRouter.POST("payPush", testJRApi.PayPush)
		testJRAPIRouter.POST("queryInfo", testJRApi.QueryInfo)
		testJRAPIRouter.POST("revokePush", testJRApi.RevokePush)
		testJRAPIRouter.POST("applyDelay", testJRApi.ApplyDelay)
		testJRAPIRouter.POST("applyRefund", testJRApi.ApplyRefund)
		testJRAPIRouter.POST("applyClaim", testJRApi.ApplyClaim)
		testJRAPIRouter.POST("logoutPush", testJRApi.LogoutPush)
		testJRAPIRouter.POST("applyInvoice", testJRApi.ApplyInvoice)
	}
	{
		testJRAPIRouterPublic.GET("letterFileDownload", testJRApi.LetterFileDownload)
		testJRAPIRouterPublic.GET("delayFileDownload", testJRApi.DelayFileDownload)
		testJRAPIRouterPublic.GET("invoiceFileDownload", testJRApi.InvoiceFileDownload)
	}
}
