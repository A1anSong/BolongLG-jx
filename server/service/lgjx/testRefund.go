package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestRefundService struct {
}

func (testRefundService *TestRefundService) CreateRefund(refund lgjx.Refund) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&refund).Error
	return err
}

func (testRefundService *TestRefundService) DeleteRefund(refund lgjx.Refund) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&refund).Error
	return err
}

func (testRefundService *TestRefundService) DeleteRefundByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Refund{}, "id in ?", ids.Ids).Error
	return err
}

func (testRefundService *TestRefundService) UpdateRefund(refund lgjx.Refund) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&refund).Error
	return err
}

func (testRefundService *TestRefundService) GetRefund(id uint) (refund lgjx.Refund, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&refund).Error
	return
}

func (testRefundService *TestRefundService) GetRefundInfoList(info lgjxReq.RefundSearch) (list []lgjx.Refund, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Refund{})
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
