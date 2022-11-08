import service from '@/utils/request'

// @Tags Invoice
// @Summary 创建Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Invoice true "创建Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoice/createInvoice [post]
export const createInvoice = (data) => {
  return service({
    url: '/invoice/createInvoice',
    method: 'post',
    data
  })
}

// @Tags Invoice
// @Summary 删除Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Invoice true "删除Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invoice/deleteInvoice [delete]
export const deleteInvoice = (data) => {
  return service({
    url: '/invoice/deleteInvoice',
    method: 'delete',
    data
  })
}

// @Tags Invoice
// @Summary 删除Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /invoice/deleteInvoice [delete]
export const deleteInvoiceByIds = (data) => {
  return service({
    url: '/invoice/deleteInvoiceByIds',
    method: 'delete',
    data
  })
}

// @Tags Invoice
// @Summary 更新Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Invoice true "更新Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /invoice/updateInvoice [put]
export const updateInvoice = (data) => {
  return service({
    url: '/invoice/updateInvoice',
    method: 'put',
    data
  })
}

// @Tags Invoice
// @Summary 用id查询Invoice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Invoice true "用id查询Invoice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /invoice/findInvoice [get]
export const findInvoice = (params) => {
  return service({
    url: '/invoice/findInvoice',
    method: 'get',
    params
  })
}

// @Tags Invoice
// @Summary 分页获取Invoice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Invoice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /invoice/getInvoiceList [get]
export const getInvoiceList = (params) => {
  return service({
    url: '/invoice/getInvoiceList',
    method: 'get',
    params
  })
}
