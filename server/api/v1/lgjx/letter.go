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

type LetterApi struct {
}

var letterService = service.ServiceGroupApp.LgjxServiceGroup.LetterService

// CreateLetter 创建Letter
// @Tags Letter
// @Summary 创建Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Letter true "创建Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letter/createLetter [post]
func (letterApi *LetterApi) CreateLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.CreateLetter(letter); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLetter 删除Letter
// @Tags Letter
// @Summary 删除Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Letter true "删除Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /letter/deleteLetter [delete]
func (letterApi *LetterApi) DeleteLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.DeleteLetter(letter); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLetterByIds 批量删除Letter
// @Tags Letter
// @Summary 批量删除Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /letter/deleteLetterByIds [delete]
func (letterApi *LetterApi) DeleteLetterByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.DeleteLetterByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLetter 更新Letter
// @Tags Letter
// @Summary 更新Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Letter true "更新Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /letter/updateLetter [put]
func (letterApi *LetterApi) UpdateLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindJSON(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := letterService.UpdateLetter(letter); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLetter 用id查询Letter
// @Tags Letter
// @Summary 用id查询Letter
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Letter true "用id查询Letter"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /letter/findLetter [get]
func (letterApi *LetterApi) FindLetter(c *gin.Context) {
	var letter lgjx.Letter
	err := c.ShouldBindQuery(&letter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reletter, err := letterService.GetLetter(letter.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reletter": reletter}, c)
	}
}

// GetLetterList 分页获取Letter列表
// @Tags Letter
// @Summary 分页获取Letter列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.LetterSearch true "分页获取Letter列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letter/getLetterList [get]
func (letterApi *LetterApi) GetLetterList(c *gin.Context) {
	var pageInfo lgjxReq.LetterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := letterService.GetLetterInfoList(pageInfo); err != nil {
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
