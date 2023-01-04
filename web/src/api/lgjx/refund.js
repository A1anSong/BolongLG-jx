import service from '@/utils/request'

// @Tags Refund
// @Summary 创建Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Refund true "创建Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /refund/createRefund [post]
export const createRefund = (data) => {
  return service({
    url: '/refund/createRefund',
    method: 'post',
    data
  })
}

// @Tags Refund
// @Summary 删除Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Refund true "删除Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /refund/deleteRefund [delete]
export const deleteRefund = (data) => {
  return service({
    url: '/refund/deleteRefund',
    method: 'delete',
    data
  })
}

// @Tags Refund
// @Summary 删除Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /refund/deleteRefund [delete]
export const deleteRefundByIds = (data) => {
  return service({
    url: '/refund/deleteRefundByIds',
    method: 'delete',
    data
  })
}

// @Tags Refund
// @Summary 更新Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Refund true "更新Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /refund/updateRefund [put]
export const updateRefund = (data) => {
  return service({
    url: '/refund/updateRefund',
    method: 'put',
    data
  })
}

// @Tags Refund
// @Summary 用id查询Refund
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Refund true "用id查询Refund"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /refund/findRefund [get]
export const findRefund = (params) => {
  return service({
    url: '/refund/findRefund',
    method: 'get',
    params
  })
}

// @Tags Refund
// @Summary 分页获取Refund列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Refund列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /refund/getRefundList [get]
export const getRefundList = (params) => {
  return service({
    url: '/refund/getRefundList',
    method: 'get',
    params
  })
}
