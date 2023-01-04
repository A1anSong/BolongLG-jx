import service from '@/utils/request'

// @Tags Delay
// @Summary 创建Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Delay true "创建Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /delay/createDelay [post]
export const createDelay = (data) => {
  return service({
    url: '/delay/createDelay',
    method: 'post',
    data
  })
}

// @Tags Delay
// @Summary 删除Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Delay true "删除Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /delay/deleteDelay [delete]
export const deleteDelay = (data) => {
  return service({
    url: '/delay/deleteDelay',
    method: 'delete',
    data
  })
}

// @Tags Delay
// @Summary 删除Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /delay/deleteDelay [delete]
export const deleteDelayByIds = (data) => {
  return service({
    url: '/delay/deleteDelayByIds',
    method: 'delete',
    data
  })
}

// @Tags Delay
// @Summary 更新Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Delay true "更新Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /delay/updateDelay [put]
export const updateDelay = (data) => {
  return service({
    url: '/delay/updateDelay',
    method: 'put',
    data
  })
}

// @Tags Delay
// @Summary 用id查询Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Delay true "用id查询Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /delay/findDelay [get]
export const findDelay = (params) => {
  return service({
    url: '/delay/findDelay',
    method: 'get',
    params
  })
}

// @Tags Delay
// @Summary 分页获取Delay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Delay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /delay/getDelayList [get]
export const getDelayList = (params) => {
  return service({
    url: '/delay/getDelayList',
    method: 'get',
    params
  })
}
