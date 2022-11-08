package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestOrderService struct {
}

func (testOrderService *TestOrderService) CreateOrder(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&order).Error
	return err
}

func (testOrderService *TestOrderService) DeleteOrder(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&order).Error
	return err
}

func (testOrderService *TestOrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Order{}, "id in ?", ids.Ids).Error
	return err
}

func (testOrderService *TestOrderService) UpdateOrder(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&order).Error
	return err
}

func (testOrderService *TestOrderService) GetOrder(id uint) (order lgjx.Order, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&order).Error
	return
}

func (testOrderService *TestOrderService) GetOrderInfoList(info lgjxReq.OrderSearch) (list []lgjx.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Order{})
	var orders []lgjx.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&orders).Error
	return orders, total, err
}
