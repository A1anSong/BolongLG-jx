import service from '@/utils/request'

// @Tags Claim
// @Summary 创建Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Claim true "创建Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /claim/createClaim [post]
export const createClaim = (data) => {
  return service({
    url: '/claim/createClaim',
    method: 'post',
    data
  })
}

// @Tags Claim
// @Summary 删除Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Claim true "删除Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /claim/deleteClaim [delete]
export const deleteClaim = (data) => {
  return service({
    url: '/claim/deleteClaim',
    method: 'delete',
    data
  })
}

// @Tags Claim
// @Summary 删除Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /claim/deleteClaim [delete]
export const deleteClaimByIds = (data) => {
  return service({
    url: '/claim/deleteClaimByIds',
    method: 'delete',
    data
  })
}

// @Tags Claim
// @Summary 更新Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Claim true "更新Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /claim/updateClaim [put]
export const updateClaim = (data) => {
  return service({
    url: '/claim/updateClaim',
    method: 'put',
    data
  })
}

// @Tags Claim
// @Summary 用id查询Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Claim true "用id查询Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /claim/findClaim [get]
export const findClaim = (params) => {
  return service({
    url: '/claim/findClaim',
    method: 'get',
    params
  })
}

// @Tags Claim
// @Summary 分页获取Claim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Claim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /claim/getClaimList [get]
export const getClaimList = (params) => {
  return service({
    url: '/claim/getClaimList',
    method: 'get',
    params
  })
}
