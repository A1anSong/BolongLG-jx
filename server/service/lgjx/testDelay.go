package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestDelayService struct {
}

func (testDelayService *TestDelayService) CreateDelay(delay lgjx.Delay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&delay).Error
	return err
}

func (testDelayService *TestDelayService) DeleteDelay(delay lgjx.Delay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&delay).Error
	return err
}

func (testDelayService *TestDelayService) DeleteDelayByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Delay{}, "id in ?", ids.Ids).Error
	return err
}

func (testDelayService *TestDelayService) UpdateDelay(delay lgjx.Delay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&delay).Error
	return err
}

func (testDelayService *TestDelayService) GetDelay(id uint) (delay lgjx.Delay, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&delay).Error
	return
}

func (testDelayService *TestDelayService) GetDelayInfoList(info lgjxReq.DelaySearch) (list []lgjx.Delay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Delay{})
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
