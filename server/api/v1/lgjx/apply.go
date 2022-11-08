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

type ApplyApi struct {
}

var applyService = service.ServiceGroupApp.LgjxServiceGroup.ApplyService

// CreateApply 创建Apply
// @Tags Apply
// @Summary 创建Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Apply true "创建Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apply/createApply [post]
func (applyApi *ApplyApi) CreateApply(c *gin.Context) {
	var apply lgjx.Apply
	err := c.ShouldBindJSON(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.CreateApply(apply); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteApply 删除Apply
// @Tags Apply
// @Summary 删除Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Apply true "删除Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apply/deleteApply [delete]
func (applyApi *ApplyApi) DeleteApply(c *gin.Context) {
	var apply lgjx.Apply
	err := c.ShouldBindJSON(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.DeleteApply(apply); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteApplyByIds 批量删除Apply
// @Tags Apply
// @Summary 批量删除Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /apply/deleteApplyByIds [delete]
func (applyApi *ApplyApi) DeleteApplyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.DeleteApplyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateApply 更新Apply
// @Tags Apply
// @Summary 更新Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Apply true "更新Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apply/updateApply [put]
func (applyApi *ApplyApi) UpdateApply(c *gin.Context) {
	var apply lgjx.Apply
	err := c.ShouldBindJSON(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := applyService.UpdateApply(apply); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindApply 用id查询Apply
// @Tags Apply
// @Summary 用id查询Apply
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Apply true "用id查询Apply"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apply/findApply [get]
func (applyApi *ApplyApi) FindApply(c *gin.Context) {
	var apply lgjx.Apply
	err := c.ShouldBindQuery(&apply)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reapply, err := applyService.GetApply(apply.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reapply": reapply}, c)
	}
}

// GetApplyList 分页获取Apply列表
// @Tags Apply
// @Summary 分页获取Apply列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.ApplySearch true "分页获取Apply列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apply/getApplyList [get]
func (applyApi *ApplyApi) GetApplyList(c *gin.Context) {
	var pageInfo lgjxReq.ApplySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := applyService.GetApplyInfoList(pageInfo); err != nil {
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
