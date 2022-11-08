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

type TestClaimApi struct {
}

var testClaimService = service.ServiceGroupApp.LgjxTestServiceGroup.TestClaimService

func (testClaimApi *TestClaimApi) CreateClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindJSON(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testClaimService.CreateClaim(claim); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (testClaimApi *TestClaimApi) DeleteClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindJSON(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testClaimService.DeleteClaim(claim); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (testClaimApi *TestClaimApi) DeleteClaimByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testClaimService.DeleteClaimByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

func (testClaimApi *TestClaimApi) UpdateClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindJSON(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := testClaimService.UpdateClaim(claim); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (testClaimApi *TestClaimApi) FindClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindQuery(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reclaim, err := testClaimService.GetClaim(claim.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclaim": reclaim}, c)
	}
}

func (testClaimApi *TestClaimApi) GetClaimList(c *gin.Context) {
	var pageInfo lgjxReq.ClaimSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := testClaimService.GetClaimInfoList(pageInfo); err != nil {
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
