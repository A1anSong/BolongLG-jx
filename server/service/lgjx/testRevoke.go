package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestRevokeService struct {
}

func (testRevokeService *TestRevokeService) CreateRevoke(revoke lgjx.Revoke) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&revoke).Error
	return err
}

func (testRevokeService *TestRevokeService) DeleteRevoke(revoke lgjx.Revoke) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&revoke).Error
	return err
}

func (testRevokeService *TestRevokeService) DeleteRevokeByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Revoke{}, "id in ?", ids.Ids).Error
	return err
}

func (testRevokeService *TestRevokeService) UpdateRevoke(revoke lgjx.Revoke) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&revoke).Error
	return err
}

func (testRevokeService *TestRevokeService) GetRevoke(id uint) (revoke lgjx.Revoke, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&revoke).Error
	return
}

func (testRevokeService *TestRevokeService) GetRevokeInfoList(info lgjxReq.RevokeSearch) (list []lgjx.Revoke, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Revoke{})
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
