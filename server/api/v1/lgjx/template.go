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

type TemplateApi struct {
}

var templateService = service.ServiceGroupApp.LgjxServiceGroup.TemplateService

// CreateTemplate 创建Template
// @Tags Template
// @Summary 创建Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Template true "创建Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /template/createTemplate [post]
func (templateApi *TemplateApi) CreateTemplate(c *gin.Context) {
	var template lgjx.Template
	err := c.ShouldBindJSON(&template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.CreateTemplate(template); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTemplate 删除Template
// @Tags Template
// @Summary 删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Template true "删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /template/deleteTemplate [delete]
func (templateApi *TemplateApi) DeleteTemplate(c *gin.Context) {
	var template lgjx.Template
	err := c.ShouldBindJSON(&template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.DeleteTemplate(template); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTemplateByIds 批量删除Template
// @Tags Template
// @Summary 批量删除Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /template/deleteTemplateByIds [delete]
func (templateApi *TemplateApi) DeleteTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.DeleteTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTemplate 更新Template
// @Tags Template
// @Summary 更新Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Template true "更新Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /template/updateTemplate [put]
func (templateApi *TemplateApi) UpdateTemplate(c *gin.Context) {
	var template lgjx.Template
	err := c.ShouldBindJSON(&template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := templateService.UpdateTemplate(template); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTemplate 用id查询Template
// @Tags Template
// @Summary 用id查询Template
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Template true "用id查询Template"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /template/findTemplate [get]
func (templateApi *TemplateApi) FindTemplate(c *gin.Context) {
	var template lgjx.Template
	err := c.ShouldBindQuery(&template)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retemplate, err := templateService.GetTemplate(template.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retemplate": retemplate}, c)
	}
}

// GetTemplateList 分页获取Template列表
// @Tags Template
// @Summary 分页获取Template列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.TemplateSearch true "分页获取Template列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /template/getTemplateList [get]
func (templateApi *TemplateApi) GetTemplateList(c *gin.Context) {
	var pageInfo lgjxReq.TemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := templateService.GetTemplateInfoList(pageInfo); err != nil {
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
