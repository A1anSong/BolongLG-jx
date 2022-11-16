import service from '@/utils/request'

// @Tags Setting
// @Summary 创建Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Setting true "创建Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /setting/createSetting [post]
export const createSetting = (data) => {
  return service({
    url: '/setting/createSetting',
    method: 'post',
    data
  })
}

// @Tags Setting
// @Summary 删除Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Setting true "删除Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /setting/deleteSetting [delete]
export const deleteSetting = (data) => {
  return service({
    url: '/setting/deleteSetting',
    method: 'delete',
    data
  })
}

// @Tags Setting
// @Summary 删除Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /setting/deleteSetting [delete]
export const deleteSettingByIds = (data) => {
  return service({
    url: '/setting/deleteSettingByIds',
    method: 'delete',
    data
  })
}

// @Tags Setting
// @Summary 更新Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Setting true "更新Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /setting/updateSetting [put]
export const updateSetting = (data) => {
  return service({
    url: '/setting/updateSetting',
    method: 'put',
    data
  })
}

// @Tags Setting
// @Summary 用id查询Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Setting true "用id查询Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /setting/findSetting [get]
export const findSetting = (params) => {
  return service({
    url: '/setting/findSetting',
    method: 'get',
    params
  })
}

// @Tags Setting
// @Summary 分页获取Setting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Setting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /setting/getSettingList [get]
export const getSettingList = (params) => {
  return service({
    url: '/setting/getSettingList',
    method: 'get',
    params
  })
}
