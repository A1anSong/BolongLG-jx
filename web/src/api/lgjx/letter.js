import service from '@/utils/request'

// @Tags Letter
// @Summary 创建Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Letter true "创建Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letter/createLetter [post]
export const createLetter = (data) => {
  return service({
    url: '/letter/createLetter',
    method: 'post',
    data
  })
}

// @Tags Letter
// @Summary 删除Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Letter true "删除Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /letter/deleteLetter [delete]
export const deleteLetter = (data) => {
  return service({
    url: '/letter/deleteLetter',
    method: 'delete',
    data
  })
}

// @Tags Letter
// @Summary 删除Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /letter/deleteLetter [delete]
export const deleteLetterByIds = (data) => {
  return service({
    url: '/letter/deleteLetterByIds',
    method: 'delete',
    data
  })
}

// @Tags Letter
// @Summary 更新Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Letter true "更新Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /letter/updateLetter [put]
export const updateLetter = (data) => {
  return service({
    url: '/letter/updateLetter',
    method: 'put',
    data
  })
}

// @Tags Letter
// @Summary 用id查询Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Letter true "用id查询Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /letter/findLetter [get]
export const findLetter = (params) => {
  return service({
    url: '/letter/findLetter',
    method: 'get',
    params
  })
}

// @Tags Letter
// @Summary 分页获取Letter列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Letter列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letter/getLetterList [get]
export const getLetterList = (params) => {
  return service({
    url: '/letter/getLetterList',
    method: 'get',
    params
  })
}
