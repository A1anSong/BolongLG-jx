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

type SettingApi struct {
}

var settingService = service.ServiceGroupApp.LgjxServiceGroup.SettingService

// CreateSetting 创建Setting
// @Tags Setting
// @Summary 创建Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Setting true "创建Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /setting/createSetting [post]
func (settingApi *SettingApi) CreateSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindJSON(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingService.CreateSetting(setting); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSetting 删除Setting
// @Tags Setting
// @Summary 删除Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Setting true "删除Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /setting/deleteSetting [delete]
func (settingApi *SettingApi) DeleteSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindJSON(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingService.DeleteSetting(setting); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSettingByIds 批量删除Setting
// @Tags Setting
// @Summary 批量删除Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /setting/deleteSettingByIds [delete]
func (settingApi *SettingApi) DeleteSettingByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingService.DeleteSettingByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSetting 更新Setting
// @Tags Setting
// @Summary 更新Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Setting true "更新Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /setting/updateSetting [put]
func (settingApi *SettingApi) UpdateSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindJSON(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingService.UpdateSetting(setting); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSetting 用id查询Setting
// @Tags Setting
// @Summary 用id查询Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Setting true "用id查询Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /setting/findSetting [get]
func (settingApi *SettingApi) FindSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindQuery(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resetting, err := settingService.GetSetting(setting.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resetting": resetting}, c)
	}
}

// GetSettingList 分页获取Setting列表
// @Tags Setting
// @Summary 分页获取Setting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.SettingSearch true "分页获取Setting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /setting/getSettingList [get]
func (settingApi *SettingApi) GetSettingList(c *gin.Context) {
	var pageInfo lgjxReq.SettingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := settingService.GetSettingInfoList(pageInfo); err != nil {
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
