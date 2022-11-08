package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type DelayService struct {
}

// CreateDelay 创建Delay记录
// Author [piexlmax](https://github.com/piexlmax)
func (delayService *DelayService) CreateDelay(delay lgjx.Delay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&delay).Error
	return err
}

// DeleteDelay 删除Delay记录
// Author [piexlmax](https://github.com/piexlmax)
func (delayService *DelayService) DeleteDelay(delay lgjx.Delay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&delay).Error
	return err
}

// DeleteDelayByIds 批量删除Delay记录
// Author [piexlmax](https://github.com/piexlmax)
func (delayService *DelayService) DeleteDelayByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Delay{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDelay 更新Delay记录
// Author [piexlmax](https://github.com/piexlmax)
func (delayService *DelayService) UpdateDelay(delay lgjx.Delay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&delay).Error
	return err
}

// GetDelay 根据id获取Delay记录
// Author [piexlmax](https://github.com/piexlmax)
func (delayService *DelayService) GetDelay(id uint) (delay lgjx.Delay, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&delay).Error
	return
}

// GetDelayInfoList 分页获取Delay记录
// Author [piexlmax](https://github.com/piexlmax)
func (delayService *DelayService) GetDelayInfoList(info lgjxReq.DelaySearch) (list []lgjx.Delay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Delay{})
	var delays []lgjx.Delay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&delays).Error
	return delays, total, err
}
