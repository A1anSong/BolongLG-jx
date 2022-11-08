package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type PayService struct {
}

// CreatePay 创建Pay记录
// Author [piexlmax](https://github.com/piexlmax)
func (payService *PayService) CreatePay(pay lgjx.Pay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&pay).Error
	return err
}

// DeletePay 删除Pay记录
// Author [piexlmax](https://github.com/piexlmax)
func (payService *PayService) DeletePay(pay lgjx.Pay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&pay).Error
	return err
}

// DeletePayByIds 批量删除Pay记录
// Author [piexlmax](https://github.com/piexlmax)
func (payService *PayService) DeletePayByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Pay{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePay 更新Pay记录
// Author [piexlmax](https://github.com/piexlmax)
func (payService *PayService) UpdatePay(pay lgjx.Pay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&pay).Error
	return err
}

// GetPay 根据id获取Pay记录
// Author [piexlmax](https://github.com/piexlmax)
func (payService *PayService) GetPay(id uint) (pay lgjx.Pay, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&pay).Error
	return
}

// GetPayInfoList 分页获取Pay记录
// Author [piexlmax](https://github.com/piexlmax)
func (payService *PayService) GetPayInfoList(info lgjxReq.PaySearch) (list []lgjx.Pay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Pay{})
	var pays []lgjx.Pay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&pays).Error
	return pays, total, err
}
