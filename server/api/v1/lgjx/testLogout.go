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

type TestLogoutApi struct {
}

var testLogoutService = service.ServiceGroupApp.LgjxTestServiceGroup.TestLogoutService

func (testLogoutApi *TestLogoutApi) CreateLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLogoutService.CreateLogout(logout); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testLogoutApi *TestLogoutApi) DeleteLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLogoutService.DeleteLogout(logout); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testLogoutApi *TestLogoutApi) DeleteLogoutByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLogoutService.DeleteLogoutByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testLogoutApi *TestLogoutApi) UpdateLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindJSON(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLogoutService.UpdateLogout(logout); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testLogoutApi *TestLogoutApi) FindLogout(c *gin.Context) {
	var logout lgjx.Logout
	err := c.ShouldBindQuery(&logout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relogout, err := testLogoutService.GetLogout(logout.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relogout": relogout}, c)
	}
}

func (testLogoutApi *TestLogoutApi) GetLogoutList(c *gin.Context) {
	var pageInfo lgjxReq.LogoutSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testLogoutService.GetLogoutInfoList(pageInfo); err != nil {
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
