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

type InvoiceApi struct {
}

var invoiceService = service.ServiceGroupApp.LgjxServiceGroup.InvoiceService

// CreateInvoice 创建Invoice
// @Tags Invoice
// @Summary 创建Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Invoice true "创建Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoice/createInvoice [post]
func (invoiceApi *InvoiceApi) CreateInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.CreateInvoice(invoice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInvoice 删除Invoice
// @Tags Invoice
// @Summary 删除Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Invoice true "删除Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invoice/deleteInvoice [delete]
func (invoiceApi *InvoiceApi) DeleteInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.DeleteInvoice(invoice); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInvoiceByIds 批量删除Invoice
// @Tags Invoice
// @Summary 批量删除Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /invoice/deleteInvoiceByIds [delete]
func (invoiceApi *InvoiceApi) DeleteInvoiceByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.DeleteInvoiceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInvoice 更新Invoice
// @Tags Invoice
// @Summary 更新Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Invoice true "更新Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /invoice/updateInvoice [put]
func (invoiceApi *InvoiceApi) UpdateInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := invoiceService.UpdateInvoice(invoice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInvoice 用id查询Invoice
// @Tags Invoice
// @Summary 用id查询Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Invoice true "用id查询Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /invoice/findInvoice [get]
func (invoiceApi *InvoiceApi) FindInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindQuery(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinvoice, err := invoiceService.GetInvoice(invoice.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinvoice": reinvoice}, c)
	}
}

// GetInvoiceList 分页获取Invoice列表
// @Tags Invoice
// @Summary 分页获取Invoice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.InvoiceSearch true "分页获取Invoice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoice/getInvoiceList [get]
func (invoiceApi *InvoiceApi) GetInvoiceList(c *gin.Context) {
	var pageInfo lgjxReq.InvoiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := invoiceService.GetInvoiceInfoList(pageInfo); err != nil {
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
