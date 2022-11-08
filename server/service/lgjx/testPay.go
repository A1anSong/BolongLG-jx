package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestPayService struct {
}

func (testPayService *TestPayService) CreatePay(pay lgjx.Pay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&pay).Error
	return err
}

func (testPayService *TestPayService) DeletePay(pay lgjx.Pay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&pay).Error
	return err
}

func (testPayService *TestPayService) DeletePayByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Pay{}, "id in ?", ids.Ids).Error
	return err
}

func (testPayService *TestPayService) UpdatePay(pay lgjx.Pay) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&pay).Error
	return err
}

func (testPayService *TestPayService) GetPay(id uint) (pay lgjx.Pay, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&pay).Error
	return
}

func (testPayService *TestPayService) GetPayInfoList(info lgjxReq.PaySearch) (list []lgjx.Pay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Pay{})
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
