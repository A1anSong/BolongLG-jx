package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestLogoutService struct {
}

func (testLogoutService *TestLogoutService) CreateLogout(logout lgjx.Logout) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&logout).Error
	return err
}

func (testLogoutService *TestLogoutService) DeleteLogout(logout lgjx.Logout) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&logout).Error
	return err
}

func (testLogoutService *TestLogoutService) DeleteLogoutByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Logout{}, "id in ?", ids.Ids).Error
	return err
}

func (testLogoutService *TestLogoutService) UpdateLogout(logout lgjx.Logout) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&logout).Error
	return err
}

func (testLogoutService *TestLogoutService) GetLogout(id uint) (logout lgjx.Logout, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&logout).Error
	return
}

func (testLogoutService *TestLogoutService) GetLogoutInfoList(info lgjxReq.LogoutSearch) (list []lgjx.Logout, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Logout{})
	var logouts []lgjx.Logout
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&logouts).Error
	return logouts, total, err
}
