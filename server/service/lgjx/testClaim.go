package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestClaimService struct {
}

func (testClaimService *TestClaimService) CreateClaim(claim lgjx.Claim) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&claim).Error
	return err
}

func (testClaimService *TestClaimService) DeleteClaim(claim lgjx.Claim) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&claim).Error
	return err
}

func (testClaimService *TestClaimService) DeleteClaimByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Claim{}, "id in ?", ids.Ids).Error
	return err
}

func (testClaimService *TestClaimService) UpdateClaim(claim lgjx.Claim) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&claim).Error
	return err
}

func (testClaimService *TestClaimService) GetClaim(id uint) (claim lgjx.Claim, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&claim).Error
	return
}

func (testClaimService *TestClaimService) GetClaimInfoList(info lgjxReq.ClaimSearch) (list []lgjx.Claim, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Claim{})
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
