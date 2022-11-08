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

type DelayApi struct {
}

var delayService = service.ServiceGroupApp.LgjxServiceGroup.DelayService

// CreateDelay 创建Delay
// @Tags Delay
// @Summary 创建Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Delay true "创建Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /delay/createDelay [post]
func (delayApi *DelayApi) CreateDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.CreateDelay(delay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDelay 删除Delay
// @Tags Delay
// @Summary 删除Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Delay true "删除Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /delay/deleteDelay [delete]
func (delayApi *DelayApi) DeleteDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.DeleteDelay(delay); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDelayByIds 批量删除Delay
// @Tags Delay
// @Summary 批量删除Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /delay/deleteDelayByIds [delete]
func (delayApi *DelayApi) DeleteDelayByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.DeleteDelayByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDelay 更新Delay
// @Tags Delay
// @Summary 更新Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Delay true "更新Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /delay/updateDelay [put]
func (delayApi *DelayApi) UpdateDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindJSON(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := delayService.UpdateDelay(delay); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDelay 用id查询Delay
// @Tags Delay
// @Summary 用id查询Delay
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Delay true "用id查询Delay"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /delay/findDelay [get]
func (delayApi *DelayApi) FindDelay(c *gin.Context) {
	var delay lgjx.Delay
	err := c.ShouldBindQuery(&delay)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redelay, err := delayService.GetDelay(delay.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redelay": redelay}, c)
	}
}

// GetDelayList 分页获取Delay列表
// @Tags Delay
// @Summary 分页获取Delay列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.DelaySearch true "分页获取Delay列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /delay/getDelayList [get]
func (delayApi *DelayApi) GetDelayList(c *gin.Context) {
	var pageInfo lgjxReq.DelaySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := delayService.GetDelayInfoList(pageInfo); err != nil {
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
