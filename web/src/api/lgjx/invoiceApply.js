import service from '@/utils/request'

// @Tags InvoiceApply
// @Summary 创建InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InvoiceApply true "创建InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoiceApply/createInvoiceApply [post]
export const createInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/createInvoiceApply',
    method: 'post',
    data
  })
}

// @Tags InvoiceApply
// @Summary 删除InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InvoiceApply true "删除InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invoiceApply/deleteInvoiceApply [delete]
export const deleteInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/deleteInvoiceApply',
    method: 'delete',
    data
  })
}

// @Tags InvoiceApply
// @Summary 删除InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invoiceApply/deleteInvoiceApply [delete]
export const deleteInvoiceApplyByIds = (data) => {
  return service({
    url: '/invoiceApply/deleteInvoiceApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags InvoiceApply
// @Summary 更新InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InvoiceApply true "更新InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /invoiceApply/updateInvoiceApply [put]
export const updateInvoiceApply = (data) => {
  return service({
    url: '/invoiceApply/updateInvoiceApply',
    method: 'put',
    data
  })
}

// @Tags InvoiceApply
// @Summary 用id查询InvoiceApply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InvoiceApply true "用id查询InvoiceApply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /invoiceApply/findInvoiceApply [get]
export const findInvoiceApply = (params) => {
  return service({
    url: '/invoiceApply/findInvoiceApply',
    method: 'get',
    params
  })
}

// @Tags InvoiceApply
// @Summary 分页获取InvoiceApply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取InvoiceApply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoiceApply/getInvoiceApplyList [get]
export const getInvoiceApplyList = (params) => {
  return service({
    url: '/invoiceApply/getInvoiceApplyList',
    method: 'get',
    params
  })
}
