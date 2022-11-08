package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestPayApi struct {
}

var TestPayService = service.ServiceGroupApp.LgjxTestServiceGroup.TestPayService

func (testPayApi *TestPayApi) CreatePay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestPayService.CreatePay(pay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testPayApi *TestPayApi) DeletePay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestPayService.DeletePay(pay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testPayApi *TestPayApi) DeletePayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestPayService.DeletePayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testPayApi *TestPayApi) UpdatePay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := TestPayService.UpdatePay(pay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testPayApi *TestPayApi) FindPay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindQuery(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repay, err := TestPayService.GetPay(pay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repay": repay}, c)
	}
}

func (testPayApi *TestPayApi) GetPayList(c *gin.Context) {
	var pageInfo lgjxReq.PaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := TestPayService.GetPayInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
