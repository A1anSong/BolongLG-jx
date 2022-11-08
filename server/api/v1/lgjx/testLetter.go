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

type TestLetterApi struct {
}

var testLetterService = service.ServiceGroupApp.LgjxTestServiceGroup.TestLetterService

func (testLetterApi *TestLetterApi) CreateLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLetterService.CreateLetter(letter); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testLetterApi *TestLetterApi) DeleteLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLetterService.DeleteLetter(letter); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testLetterApi *TestLetterApi) DeleteLetterByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLetterService.DeleteLetterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testLetterApi *TestLetterApi) UpdateLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testLetterService.UpdateLetter(letter); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testLetterApi *TestLetterApi) FindLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindQuery(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reletter, err := testLetterService.GetLetter(letter.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reletter": reletter}, c)
	}
}

func (testLetterApi *TestLetterApi) GetLetterList(c *gin.Context) {
	var pageInfo lgjxReq.LetterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testLetterService.GetLetterInfoList(pageInfo); err != nil {
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
