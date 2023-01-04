import service from '@/utils/request'

// @Tags Apply
// @Summary 创建Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apply true "创建Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apply/createApply [post]
export const createApply = (data) => {
  return service({
    url: '/apply/createApply',
    method: 'post',
    data
  })
}

// @Tags Apply
// @Summary 删除Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apply true "删除Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apply/deleteApply [delete]
export const deleteApply = (data) => {
  return service({
    url: '/apply/deleteApply',
    method: 'delete',
    data
  })
}

// @Tags Apply
// @Summary 删除Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apply/deleteApply [delete]
export const deleteApplyByIds = (data) => {
  return service({
    url: '/apply/deleteApplyByIds',
    method: 'delete',
    data
  })
}

// @Tags Apply
// @Summary 更新Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Apply true "更新Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apply/updateApply [put]
export const updateApply = (data) => {
  return service({
    url: '/apply/updateApply',
    method: 'put',
    data
  })
}

// @Tags Apply
// @Summary 用id查询Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Apply true "用id查询Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apply/findApply [get]
export const findApply = (params) => {
  return service({
    url: '/apply/findApply',
    method: 'get',
    params
  })
}

// @Tags Apply
// @Summary 分页获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Apply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apply/getApplyList [get]
export const getApplyList = (params) => {
  return service({
    url: '/apply/getApplyList',
    method: 'get',
    params
  })
}
