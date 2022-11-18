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

type TestFileApi struct {
}

var testFileService = service.ServiceGroupApp.LgjxTestServiceGroup.TestFileService

func (testFileApi *TestFileApi) CreateFile(c *gin.Context) {
	var file lgjx.File
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testFileService.CreateFile(file); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testFileApi *TestFileApi) DeleteFile(c *gin.Context) {
	var file lgjx.File
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testFileService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testFileApi *TestFileApi) DeleteFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testFileService.DeleteFileByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testFileApi *TestFileApi) UpdateFile(c *gin.Context) {
	var file lgjx.File
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testFileService.UpdateFile(file); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testFileApi *TestFileApi) FindFile(c *gin.Context) {
	var file lgjx.File
	err := c.ShouldBindQuery(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refile, err := testFileService.GetFile(file.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refile": refile}, c)
	}
}

func (testFileApi *TestFileApi) GetFileList(c *gin.Context) {
	var pageInfo lgjxReq.FileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testFileService.GetFileInfoList(pageInfo); err != nil {
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

func (testFileApi *TestFileApi) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	if fileName, err := testFileService.UploadFile(file); err != nil {
		global.GVA_LOG.Error("写入文件失败!", zap.Error(err))
		response.FailWithMessage("写入文件失败", c)
		return
	} else {
		response.OkWithDetailed(gin.H{
			"fileName": fileName,
		}, "获取成功", c)
	}
}
