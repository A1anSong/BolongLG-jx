package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestSettingService struct {
}

func (testSettingService *TestSettingService) CreateSetting(setting lgjx.Setting) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&setting).Error
	return err
}

func (testSettingService *TestSettingService) DeleteSetting(setting lgjx.Setting) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&setting).Error
	return err
}

func (testSettingService *TestSettingService) DeleteSettingByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Setting{}, "id in ?", ids.Ids).Error
	return err
}

func (testSettingService *TestSettingService) UpdateSetting(setting lgjx.Setting) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&setting).Error
	return err
}

func (testSettingService *TestSettingService) GetSetting(id uint) (setting lgjx.Setting, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&setting).Error
	return
}

func (testSettingService *TestSettingService) GetSettingInfoList(info lgjxReq.SettingSearch) (list []lgjx.Setting, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Setting{})
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
