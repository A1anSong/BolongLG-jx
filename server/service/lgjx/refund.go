package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type RefundService struct {
}

// CreateRefund 创建Refund记录
// Author [piexlmax](https://github.com/piexlmax)
func (refundService *RefundService) CreateRefund(refund lgjx.Refund) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&refund).Error
	return err
}

// DeleteRefund 删除Refund记录
// Author [piexlmax](https://github.com/piexlmax)
func (refundService *RefundService) DeleteRefund(refund lgjx.Refund) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&refund).Error
	return err
}

// DeleteRefundByIds 批量删除Refund记录
// Author [piexlmax](https://github.com/piexlmax)
func (refundService *RefundService) DeleteRefundByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Refund{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRefund 更新Refund记录
// Author [piexlmax](https://github.com/piexlmax)
func (refundService *RefundService) UpdateRefund(refund lgjx.Refund) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&refund).Error
	return err
}

// GetRefund 根据id获取Refund记录
// Author [piexlmax](https://github.com/piexlmax)
func (refundService *RefundService) GetRefund(id uint) (refund lgjx.Refund, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&refund).Error
	return
}

// GetRefundInfoList 分页获取Refund记录
// Author [piexlmax](https://github.com/piexlmax)
func (refundService *RefundService) GetRefundInfoList(info lgjxReq.RefundSearch) (list []lgjx.Refund, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Refund{})
	var refunds []lgjx.Refund
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&refunds).Error
	return refunds, total, err
}
