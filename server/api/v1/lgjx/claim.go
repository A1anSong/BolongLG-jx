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

type ClaimApi struct {
}

var claimService = service.ServiceGroupApp.LgjxServiceGroup.ClaimService

// CreateClaim 创建Claim
// @Tags Claim
// @Summary 创建Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Claim true "创建Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /claim/createClaim [post]
func (claimApi *ClaimApi) CreateClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindJSON(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := claimService.CreateClaim(claim); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClaim 删除Claim
// @Tags Claim
// @Summary 删除Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Claim true "删除Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /claim/deleteClaim [delete]
func (claimApi *ClaimApi) DeleteClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindJSON(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := claimService.DeleteClaim(claim); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClaimByIds 批量删除Claim
// @Tags Claim
// @Summary 批量删除Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /claim/deleteClaimByIds [delete]
func (claimApi *ClaimApi) DeleteClaimByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := claimService.DeleteClaimByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClaim 更新Claim
// @Tags Claim
// @Summary 更新Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body lgjx.Claim true "更新Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /claim/updateClaim [put]
func (claimApi *ClaimApi) UpdateClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindJSON(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := claimService.UpdateClaim(claim); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClaim 用id查询Claim
// @Tags Claim
// @Summary 用id查询Claim
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjx.Claim true "用id查询Claim"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /claim/findClaim [get]
func (claimApi *ClaimApi) FindClaim(c *gin.Context) {
	var claim lgjx.Claim
	err := c.ShouldBindQuery(&claim)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reclaim, err := claimService.GetClaim(claim.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclaim": reclaim}, c)
	}
}

// GetClaimList 分页获取Claim列表
// @Tags Claim
// @Summary 分页获取Claim列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query lgjxReq.ClaimSearch true "分页获取Claim列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /claim/getClaimList [get]
func (claimApi *ClaimApi) GetClaimList(c *gin.Context) {
	var pageInfo lgjxReq.ClaimSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := claimService.GetClaimInfoList(pageInfo); err != nil {
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
