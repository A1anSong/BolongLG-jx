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

type TestSettingApi struct {
}

var testSettingService = service.ServiceGroupApp.LgjxTestServiceGroup.TestSettingService

func (testSettingApi *TestSettingApi) CreateSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindJSON(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testSettingService.CreateSetting(setting); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testSettingApi *TestSettingApi) DeleteSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindJSON(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testSettingService.DeleteSetting(setting); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testSettingApi *TestSettingApi) DeleteSettingByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testSettingService.DeleteSettingByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testSettingApi *TestSettingApi) UpdateSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindJSON(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testSettingService.UpdateSetting(setting); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testSettingApi *TestSettingApi) FindSetting(c *gin.Context) {
	var setting lgjx.Setting
	err := c.ShouldBindQuery(&setting)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resetting, err := testSettingService.GetSetting(setting.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resetting": resetting}, c)
	}
}

func (testSettingApi *TestSettingApi) GetSettingList(c *gin.Context) {
	var pageInfo lgjxReq.SettingSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testSettingService.GetSettingInfoList(pageInfo); err != nil {
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
