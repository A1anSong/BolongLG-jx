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

type TestRevokeApi struct {
}

var testRevokeService = service.ServiceGroupApp.LgjxTestServiceGroup.TestRevokeService

func (testRevokeApi *TestRevokeApi) CreateRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testRevokeService.CreateRevoke(revoke); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testRevokeApi *TestRevokeApi) DeleteRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testRevokeService.DeleteRevoke(revoke); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testRevokeApi *TestRevokeApi) DeleteRevokeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testRevokeService.DeleteRevokeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testRevokeApi *TestRevokeApi) UpdateRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindJSON(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testRevokeService.UpdateRevoke(revoke); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testRevokeApi *TestRevokeApi) FindRevoke(c *gin.Context) {
	var revoke lgjx.Revoke
	err := c.ShouldBindQuery(&revoke)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerevoke, err := testRevokeService.GetRevoke(revoke.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerevoke": rerevoke}, c)
	}
}

func (testRevokeApi *TestRevokeApi) GetRevokeList(c *gin.Context) {
	var pageInfo lgjxReq.RevokeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testRevokeService.GetRevokeInfoList(pageInfo); err != nil {
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
