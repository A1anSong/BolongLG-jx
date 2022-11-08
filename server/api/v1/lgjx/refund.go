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

type RefundApi struct {
}

var refundService = service.ServiceGroupApp.LgjxServiceGroup.RefundService

// CreateRefund 创建Refund
// @Tags Refund
// @Summary 创建Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Refund true "创建Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /refund/createRefund [post]
func (refundApi *RefundApi) CreateRefund(c *gin.Context) {
	var refund lgjx.Refund
	err := c.ShouldBindJSON(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.CreateRefund(refund); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRefund 删除Refund
// @Tags Refund
// @Summary 删除Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Refund true "删除Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /refund/deleteRefund [delete]
func (refundApi *RefundApi) DeleteRefund(c *gin.Context) {
	var refund lgjx.Refund
	err := c.ShouldBindJSON(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.DeleteRefund(refund); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRefundByIds 批量删除Refund
// @Tags Refund
// @Summary 批量删除Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /refund/deleteRefundByIds [delete]
func (refundApi *RefundApi) DeleteRefundByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.DeleteRefundByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRefund 更新Refund
// @Tags Refund
// @Summary 更新Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Refund true "更新Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /refund/updateRefund [put]
func (refundApi *RefundApi) UpdateRefund(c *gin.Context) {
	var refund lgjx.Refund
	err := c.ShouldBindJSON(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := refundService.UpdateRefund(refund); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRefund 用id查询Refund
// @Tags Refund
// @Summary 用id查询Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Refund true "用id查询Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /refund/findRefund [get]
func (refundApi *RefundApi) FindRefund(c *gin.Context) {
	var refund lgjx.Refund
	err := c.ShouldBindQuery(&refund)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerefund, err := refundService.GetRefund(refund.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerefund": rerefund}, c)
	}
}

// GetRefundList 分页获取Refund列表
// @Tags Refund
// @Summary 分页获取Refund列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.RefundSearch true "分页获取Refund列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /refund/getRefundList [get]
func (refundApi *RefundApi) GetRefundList(c *gin.Context) {
	var pageInfo lgjxReq.RefundSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := refundService.GetRefundInfoList(pageInfo); err != nil {
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
