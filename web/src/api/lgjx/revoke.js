import service from '@/utils/request'

// @Tags Revoke
// @Summary 创建Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Revoke true "创建Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /revoke/createRevoke [post]
export const createRevoke = (data) => {
  return service({
    url: '/revoke/createRevoke',
    method: 'post',
    data
  })
}

// @Tags Revoke
// @Summary 删除Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Revoke true "删除Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /revoke/deleteRevoke [delete]
export const deleteRevoke = (data) => {
  return service({
    url: '/revoke/deleteRevoke',
    method: 'delete',
    data
  })
}

// @Tags Revoke
// @Summary 删除Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /revoke/deleteRevoke [delete]
export const deleteRevokeByIds = (data) => {
  return service({
    url: '/revoke/deleteRevokeByIds',
    method: 'delete',
    data
  })
}

// @Tags Revoke
// @Summary 更新Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Revoke true "更新Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /revoke/updateRevoke [put]
export const updateRevoke = (data) => {
  return service({
    url: '/revoke/updateRevoke',
    method: 'put',
    data
  })
}

// @Tags Revoke
// @Summary 用id查询Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Revoke true "用id查询Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /revoke/findRevoke [get]
export const findRevoke = (params) => {
  return service({
    url: '/revoke/findRevoke',
    method: 'get',
    params
  })
}

// @Tags Revoke
// @Summary 分页获取Revoke列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Revoke列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /revoke/getRevokeList [get]
export const getRevokeList = (params) => {
  return service({
    url: '/revoke/getRevokeList',
    method: 'get',
    params
  })
}
