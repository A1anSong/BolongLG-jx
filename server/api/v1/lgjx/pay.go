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

type PayApi struct {
}

var payService = service.ServiceGroupApp.LgjxServiceGroup.PayService

// CreatePay 创建Pay
// @Tags Pay
// @Summary 创建Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Pay true "创建Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pay/createPay [post]
func (payApi *PayApi) CreatePay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.CreatePay(pay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePay 删除Pay
// @Tags Pay
// @Summary 删除Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Pay true "删除Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pay/deletePay [delete]
func (payApi *PayApi) DeletePay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.DeletePay(pay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayByIds 批量删除Pay
// @Tags Pay
// @Summary 批量删除Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pay/deletePayByIds [delete]
func (payApi *PayApi) DeletePayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.DeletePayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePay 更新Pay
// @Tags Pay
// @Summary 更新Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Pay true "更新Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pay/updatePay [put]
func (payApi *PayApi) UpdatePay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindJSON(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payService.UpdatePay(pay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPay 用id查询Pay
// @Tags Pay
// @Summary 用id查询Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Pay true "用id查询Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pay/findPay [get]
func (payApi *PayApi) FindPay(c *gin.Context) {
	var pay lgjx.Pay
	err := c.ShouldBindQuery(&pay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repay, err := payService.GetPay(pay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repay": repay}, c)
	}
}

// GetPayList 分页获取Pay列表
// @Tags Pay
// @Summary 分页获取Pay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.PaySearch true "分页获取Pay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pay/getPayList [get]
func (payApi *PayApi) GetPayList(c *gin.Context) {
	var pageInfo lgjxReq.PaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := payService.GetPayInfoList(pageInfo); err != nil {
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
