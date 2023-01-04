import service from '@/utils/request'

// @Tags Logout
// @Summary 创建Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Logout true "创建Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logout/createLogout [post]
export const createLogout = (data) => {
  return service({
    url: '/logout/createLogout',
    method: 'post',
    data
  })
}

// @Tags Logout
// @Summary 删除Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Logout true "删除Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logout/deleteLogout [delete]
export const deleteLogout = (data) => {
  return service({
    url: '/logout/deleteLogout',
    method: 'delete',
    data
  })
}

// @Tags Logout
// @Summary 删除Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logout/deleteLogout [delete]
export const deleteLogoutByIds = (data) => {
  return service({
    url: '/logout/deleteLogoutByIds',
    method: 'delete',
    data
  })
}

// @Tags Logout
// @Summary 更新Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Logout true "更新Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /logout/updateLogout [put]
export const updateLogout = (data) => {
  return service({
    url: '/logout/updateLogout',
    method: 'put',
    data
  })
}

// @Tags Logout
// @Summary 用id查询Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Logout true "用id查询Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /logout/findLogout [get]
export const findLogout = (params) => {
  return service({
    url: '/logout/findLogout',
    method: 'get',
    params
  })
}

// @Tags Logout
// @Summary 分页获取Logout列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Logout列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logout/getLogoutList [get]
export const getLogoutList = (params) => {
  return service({
    url: '/logout/getLogoutList',
    method: 'get',
    params
  })
}
