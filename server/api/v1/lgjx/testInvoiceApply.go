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

type TestInvoiceApplyApi struct {
}

var testInvoiceApplyService = service.ServiceGroupApp.LgjxTestServiceGroup.TestInvoiceApplyService

func (testInvoiceApplyApi *TestInvoiceApplyApi) CreateInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceApplyService.CreateInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testInvoiceApplyApi *TestInvoiceApplyApi) DeleteInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceApplyService.DeleteInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testInvoiceApplyApi *TestInvoiceApplyApi) DeleteInvoiceApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceApplyService.DeleteInvoiceApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testInvoiceApplyApi *TestInvoiceApplyApi) UpdateInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindJSON(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testInvoiceApplyService.UpdateInvoiceApply(invoiceApply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testInvoiceApplyApi *TestInvoiceApplyApi) FindInvoiceApply(c *gin.Context) {
	var invoiceApply lgjx.InvoiceApply
	err := c.ShouldBindQuery(&invoiceApply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinvoiceApply, err := testInvoiceApplyService.GetInvoiceApply(invoiceApply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinvoiceApply": reinvoiceApply}, c)
	}
}

func (testInvoiceApplyApi *TestInvoiceApplyApi) GetInvoiceApplyList(c *gin.Context) {
	var pageInfo lgjxReq.InvoiceApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testInvoiceApplyService.GetInvoiceApplyInfoList(pageInfo); err != nil {
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
