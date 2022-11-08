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

type LogoutApi struct {
}

var logoutService = service.ServiceGroupApp.LgjxServiceGroup.LogoutService

// CreateLogout 创建Logout
// @Tags Logout
// @Summary 创建Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Logout true "创建Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logout/createLogout [post]
func (logoutApi *LogoutApi) CreateLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.CreateLogout(logout); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLogout 删除Logout
// @Tags Logout
// @Summary 删除Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Logout true "删除Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /logout/deleteLogout [delete]
func (logoutApi *LogoutApi) DeleteLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.DeleteLogout(logout); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLogoutByIds 批量删除Logout
// @Tags Logout
// @Summary 批量删除Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /logout/deleteLogoutByIds [delete]
func (logoutApi *LogoutApi) DeleteLogoutByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.DeleteLogoutByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLogout 更新Logout
// @Tags Logout
// @Summary 更新Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Logout true "更新Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /logout/updateLogout [put]
func (logoutApi *LogoutApi) UpdateLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := logoutService.UpdateLogout(logout); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLogout 用id查询Logout
// @Tags Logout
// @Summary 用id查询Logout
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Logout true "用id查询Logout"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /logout/findLogout [get]
func (logoutApi *LogoutApi) FindLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindQuery(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relogout, err := logoutService.GetLogout(logout.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relogout": relogout}, c)
	}
}

// GetLogoutList 分页获取Logout列表
// @Tags Logout
// @Summary 分页获取Logout列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.LogoutSearch true "分页获取Logout列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /logout/getLogoutList [get]
func (logoutApi *LogoutApi) GetLogoutList(c *gin.Context) {
	var pageInfo lgjxReq.LogoutSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := logoutService.GetLogoutInfoList(pageInfo); err != nil {
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
