package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RevokeApi struct {
}

var revokeService = service.ServiceGroupApp.LgjxServiceGroup.RevokeService

// CreateRevoke 创建Revoke
// @Tags Revoke
// @Summary 创建Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Revoke true "创建Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /revoke/createRevoke [post]
func (revokeApi *RevokeApi) CreateRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.CreateRevoke(revoke); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRevoke 删除Revoke
// @Tags Revoke
// @Summary 删除Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Revoke true "删除Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /revoke/deleteRevoke [delete]
func (revokeApi *RevokeApi) DeleteRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.DeleteRevoke(revoke); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRevokeByIds 批量删除Revoke
// @Tags Revoke
// @Summary 批量删除Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /revoke/deleteRevokeByIds [delete]
func (revokeApi *RevokeApi) DeleteRevokeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.DeleteRevokeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRevoke 更新Revoke
// @Tags Revoke
// @Summary 更新Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Revoke true "更新Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /revoke/updateRevoke [put]
func (revokeApi *RevokeApi) UpdateRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := revokeService.UpdateRevoke(revoke); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRevoke 用id查询Revoke
// @Tags Revoke
// @Summary 用id查询Revoke
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Revoke true "用id查询Revoke"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /revoke/findRevoke [get]
func (revokeApi *RevokeApi) FindRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindQuery(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerevoke, err := revokeService.GetRevoke(revoke.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerevoke": rerevoke}, c)
	}
}

// GetRevokeList 分页获取Revoke列表
// @Tags Revoke
// @Summary 分页获取Revoke列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.RevokeSearch true "分页获取Revoke列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /revoke/getRevokeList [get]
func (revokeApi *RevokeApi) GetRevokeList(c *gin.Context) {
	var pageInfo lgjxReq.RevokeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := revokeService.GetRevokeInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
