package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type ClaimService struct {
}

// CreateClaim 创建Claim记录
// Author [piexlmax](https://github.com/piexlmax)
func (claimService *ClaimService) CreateClaim(claim lgjx.Claim) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&claim).Error
	return err
}

// DeleteClaim 删除Claim记录
// Author [piexlmax](https://github.com/piexlmax)
func (claimService *ClaimService) DeleteClaim(claim lgjx.Claim) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&claim).Error
	return err
}

// DeleteClaimByIds 批量删除Claim记录
// Author [piexlmax](https://github.com/piexlmax)
func (claimService *ClaimService) DeleteClaimByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Claim{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateClaim 更新Claim记录
// Author [piexlmax](https://github.com/piexlmax)
func (claimService *ClaimService) UpdateClaim(claim lgjx.Claim) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&claim).Error
	return err
}

// GetClaim 根据id获取Claim记录
// Author [piexlmax](https://github.com/piexlmax)
func (claimService *ClaimService) GetClaim(id uint) (claim lgjx.Claim, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&claim).Error
	return
}

// GetClaimInfoList 分页获取Claim记录
// Author [piexlmax](https://github.com/piexlmax)
func (claimService *ClaimService) GetClaimInfoList(info lgjxReq.ClaimSearch) (list []lgjx.Claim, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Claim{})
	var claims []lgjx.Claim
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&claims).Error
	return claims, total, err
}
