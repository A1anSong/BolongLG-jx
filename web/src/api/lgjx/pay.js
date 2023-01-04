import service from '@/utils/request'

// @Tags Pay
// @Summary 创建Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pay true "创建Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pay/createPay [post]
export const createPay = (data) => {
  return service({
    url: '/pay/createPay',
    method: 'post',
    data
  })
}

// @Tags Pay
// @Summary 删除Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pay true "删除Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pay/deletePay [delete]
export const deletePay = (data) => {
  return service({
    url: '/pay/deletePay',
    method: 'delete',
    data
  })
}

// @Tags Pay
// @Summary 删除Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pay/deletePay [delete]
export const deletePayByIds = (data) => {
  return service({
    url: '/pay/deletePayByIds',
    method: 'delete',
    data
  })
}

// @Tags Pay
// @Summary 更新Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Pay true "更新Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pay/updatePay [put]
export const updatePay = (data) => {
  return service({
    url: '/pay/updatePay',
    method: 'put',
    data
  })
}

// @Tags Pay
// @Summary 用id查询Pay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Pay true "用id查询Pay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pay/findPay [get]
export const findPay = (params) => {
  return service({
    url: '/pay/findPay',
    method: 'get',
    params
  })
}

// @Tags Pay
// @Summary 分页获取Pay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Pay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pay/getPayList [get]
export const getPayList = (params) => {
  return service({
    url: '/pay/getPayList',
    method: 'get',
    params
  })
}
