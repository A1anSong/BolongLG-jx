package lgjx

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type JRAPIRouter struct {
}

func (s *JRAPIRouter) InitJRAPIRouter(Router *gin.RouterGroup) {
	jrapiRouter := Router.Group("lg").Use(middleware.JRValidate())
	jrApi := v1.ApiGroupApp.LgjxApiGroup.JRAPI
	{
		jrapiRouter.POST("apply", jrApi.Apply)
		jrapiRouter.POST("payPush", jrApi.PayPush)
		jrapiRouter.POST("queryInfo", jrApi.QueryInfo)
		jrapiRouter.POST("revokePush", jrApi.RevokePush)
		jrapiRouter.POST("applyDelay", jrApi.ApplyDelay)
		jrapiRouter.POST("applyRefund", jrApi.ApplyRefund)
		jrapiRouter.POST("applyClaim", jrApi.ApplyClaim)
		jrapiRouter.POST("logoutPush", jrApi.LogoutPush)
		jrapiRouter.POST("applyInvoice", jrApi.ApplyInvoice)
	}
	{
		jrapiRouter.GET("letterFileDownload", jrApi.LetterFileDownload)
		jrapiRouter.GET("delayFileDownload", jrApi.DelayFileDownload)
	}
}
