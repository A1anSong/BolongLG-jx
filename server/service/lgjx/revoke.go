package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type RevokeService struct {
}

// CreateRevoke 创建Revoke记录
// Author [piexlmax](https://github.com/piexlmax)
func (revokeService *RevokeService) CreateRevoke(revoke lgjx.Revoke) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&revoke).Error
	return err
}

// DeleteRevoke 删除Revoke记录
// Author [piexlmax](https://github.com/piexlmax)
func (revokeService *RevokeService) DeleteRevoke(revoke lgjx.Revoke) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&revoke).Error
	return err
}

// DeleteRevokeByIds 批量删除Revoke记录
// Author [piexlmax](https://github.com/piexlmax)
func (revokeService *RevokeService) DeleteRevokeByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Revoke{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRevoke 更新Revoke记录
// Author [piexlmax](https://github.com/piexlmax)
func (revokeService *RevokeService) UpdateRevoke(revoke lgjx.Revoke) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&revoke).Error
	return err
}

// GetRevoke 根据id获取Revoke记录
// Author [piexlmax](https://github.com/piexlmax)
func (revokeService *RevokeService) GetRevoke(id uint) (revoke lgjx.Revoke, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&revoke).Error
	return
}

// GetRevokeInfoList 分页获取Revoke记录
// Author [piexlmax](https://github.com/piexlmax)
func (revokeService *RevokeService) GetRevokeInfoList(info lgjxReq.RevokeSearch) (list []lgjx.Revoke, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Revoke{})
	var revokes []lgjx.Revoke
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&revokes).Error
	return revokes, total, err
}
