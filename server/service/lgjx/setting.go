package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type SettingService struct {
}

// CreateSetting 创建Setting记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingService *SettingService) CreateSetting(setting lgjx.Setting) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&setting).Error
	return err
}

// DeleteSetting 删除Setting记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingService *SettingService) DeleteSetting(setting lgjx.Setting) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&setting).Error
	return err
}

// DeleteSettingByIds 批量删除Setting记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingService *SettingService) DeleteSettingByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Setting{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSetting 更新Setting记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingService *SettingService) UpdateSetting(setting lgjx.Setting) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&setting).Error
	return err
}

// GetSetting 根据id获取Setting记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingService *SettingService) GetSetting(id uint) (setting lgjx.Setting, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&setting).Error
	return
}

// GetSettingInfoList 分页获取Setting记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingService *SettingService) GetSettingInfoList(info lgjxReq.SettingSearch) (list []lgjx.Setting, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Setting{})
	var settings []lgjx.Setting
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&settings).Error
	return settings, total, err
}
