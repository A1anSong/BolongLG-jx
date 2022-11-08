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

type TestInvoiceApi struct {
}

var testInvoiceService = service.ServiceGroupApp.LgjxTestServiceGroup.TestInvoiceService

func (testInvoiceApi *TestInvoiceApi) CreateInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceService.CreateInvoice(invoice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testInvoiceApi *TestInvoiceApi) DeleteInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceService.DeleteInvoice(invoice); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testInvoiceApi *TestInvoiceApi) DeleteInvoiceByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceService.DeleteInvoiceByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testInvoiceApi *TestInvoiceApi) UpdateInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceService.UpdateInvoice(invoice); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testInvoiceApi *TestInvoiceApi) FindInvoice(c *gin.Context) {
	var invoice lgjx.Invoice
	err := c.ShouldBindQuery(&invoice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinvoice, err := testInvoiceService.GetInvoice(invoice.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinvoice": reinvoice}, c)
	}
}

func (testInvoiceApi *TestInvoiceApi) GetInvoiceList(c *gin.Context) {
	var pageInfo lgjxReq.InvoiceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testInvoiceService.GetInvoiceInfoList(pageInfo); err != nil {
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
