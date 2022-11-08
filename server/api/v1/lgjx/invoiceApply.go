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

type InvoiceApplyApi struct {
}

var invoiceApplyService = service.ServiceGroupApp.LgjxServiceGroup.InvoiceApplyService

// CreateInvoiceApply 创建InvoiceApply
// @Tags InvoiceApply
// @Summary 创建InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.InvoiceApply true "创建InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoiceApply/createInvoiceApply [post]
func (invoiceApplyApi *InvoiceApplyApi) CreateInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.CreateInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInvoiceApply 删除InvoiceApply
// @Tags InvoiceApply
// @Summary 删除InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.InvoiceApply true "删除InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invoiceApply/deleteInvoiceApply [delete]
func (invoiceApplyApi *InvoiceApplyApi) DeleteInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.DeleteInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInvoiceApplyByIds 批量删除InvoiceApply
// @Tags InvoiceApply
// @Summary 批量删除InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /invoiceApply/deleteInvoiceApplyByIds [delete]
func (invoiceApplyApi *InvoiceApplyApi) DeleteInvoiceApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.DeleteInvoiceApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInvoiceApply 更新InvoiceApply
// @Tags InvoiceApply
// @Summary 更新InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.InvoiceApply true "更新InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /invoiceApply/updateInvoiceApply [put]
func (invoiceApplyApi *InvoiceApplyApi) UpdateInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceApplyService.UpdateInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInvoiceApply 用id查询InvoiceApply
// @Tags InvoiceApply
// @Summary 用id查询InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.InvoiceApply true "用id查询InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /invoiceApply/findInvoiceApply [get]
func (invoiceApplyApi *InvoiceApplyApi) FindInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindQuery(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinvoiceApply, err := invoiceApplyService.GetInvoiceApply(invoiceApply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinvoiceApply": reinvoiceApply}, c)
	}
}

// GetInvoiceApplyList 分页获取InvoiceApply列表
// @Tags InvoiceApply
// @Summary 分页获取InvoiceApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.InvoiceApplySearch true "分页获取InvoiceApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoiceApply/getInvoiceApplyList [get]
func (invoiceApplyApi *InvoiceApplyApi) GetInvoiceApplyList(c *gin.Context) {
	var pageInfo lgjxReq.InvoiceApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := invoiceApplyService.GetInvoiceApplyInfoList(pageInfo); err != nil {
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
