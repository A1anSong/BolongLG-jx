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

type TestProjectApi struct {
}

var testProjectService = service.ServiceGroupApp.LgjxTestServiceGroup.TestProjectService

func (testProjectApi *TestProjectApi) CreateProject(c *gin.Context) {
	var project lgjx.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testProjectService.CreateProject(project); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testProjectApi *TestProjectApi) DeleteProject(c *gin.Context) {
	var project lgjx.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testProjectService.DeleteProject(project); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testProjectApi *TestProjectApi) DeleteProjectByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testProjectService.DeleteProjectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testProjectApi *TestProjectApi) UpdateProject(c *gin.Context) {
	var project lgjx.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testProjectService.UpdateProject(project); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testProjectApi *TestProjectApi) FindProject(c *gin.Context) {
	var project lgjx.Project
	err := c.ShouldBindQuery(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reproject, err := testProjectService.GetProject(project.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproject": reproject}, c)
	}
}

func (testProjectApi *TestProjectApi) GetProjectList(c *gin.Context) {
	var pageInfo lgjxReq.ProjectSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testProjectService.GetProjectInfoList(pageInfo); err != nil {
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

func (testProjectApi *TestProjectApi) BindProject(c *gin.Context) {
	var project lgjx.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testProjectService.BindProject(project); err != nil {
		global.GVA_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithMessage("绑定失败", c)
	} else {
		response.OkWithMessage("绑定成功", c)
	}
}

func (testProjectApi *TestProjectApi) UnbindProject(c *gin.Context) {
	var project lgjx.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testProjectService.UnbindProject(project); err != nil {
		global.GVA_LOG.Error("解绑失败!", zap.Error(err))
		response.FailWithMessage("解绑失败", c)
	} else {
		response.OkWithMessage("解绑成功", c)
	}
}
