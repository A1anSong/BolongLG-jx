package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type LogoutService struct {
}

// CreateLogout 创建Logout记录
// Author [piexlmax](https://github.com/piexlmax)
func (logoutService *LogoutService) CreateLogout(logout lgjx.Logout) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&logout).Error
	return err
}

// DeleteLogout 删除Logout记录
// Author [piexlmax](https://github.com/piexlmax)
func (logoutService *LogoutService) DeleteLogout(logout lgjx.Logout) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&logout).Error
	return err
}

// DeleteLogoutByIds 批量删除Logout记录
// Author [piexlmax](https://github.com/piexlmax)
func (logoutService *LogoutService) DeleteLogoutByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Logout{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateLogout 更新Logout记录
// Author [piexlmax](https://github.com/piexlmax)
func (logoutService *LogoutService) UpdateLogout(logout lgjx.Logout) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&logout).Error
	return err
}

// GetLogout 根据id获取Logout记录
// Author [piexlmax](https://github.com/piexlmax)
func (logoutService *LogoutService) GetLogout(id uint) (logout lgjx.Logout, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&logout).Error
	return
}

// GetLogoutInfoList 分页获取Logout记录
// Author [piexlmax](https://github.com/piexlmax)
func (logoutService *LogoutService) GetLogoutInfoList(info lgjxReq.LogoutSearch) (list []lgjx.Logout, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Logout{})
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
