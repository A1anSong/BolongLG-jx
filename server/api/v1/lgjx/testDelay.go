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

type TestDelayApi struct {
}

var testDelayService = service.ServiceGroupApp.LgjxTestServiceGroup.TestDelayService

func (testDelayApi *TestDelayApi) CreateDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testDelayService.CreateDelay(delay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testDelayApi *TestDelayApi) DeleteDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testDelayService.DeleteDelay(delay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testDelayApi *TestDelayApi) DeleteDelayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testDelayService.DeleteDelayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testDelayApi *TestDelayApi) UpdateDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testDelayService.UpdateDelay(delay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testDelayApi *TestDelayApi) FindDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindQuery(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redelay, err := testDelayService.GetDelay(delay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redelay": redelay}, c)
	}
}

func (testDelayApi *TestDelayApi) GetDelayList(c *gin.Context) {
	var pageInfo lgjxReq.DelaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testDelayService.GetDelayInfoList(pageInfo); err != nil {
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
