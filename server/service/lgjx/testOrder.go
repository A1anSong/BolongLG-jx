package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
	"gorm.io/gorm/clause"
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
	db.Joins("Apply").Joins("Pay").Joins("Letter").Joins("Revoke").Joins("Delay").Joins("Refund").Joins("Logout").Joins("Invoice").Joins("Project")
	var orders []lgjx.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ApplyNo != nil {
		db = db.Where("Apply.apply_no = ?", info.ApplyNo)
	}
	if info.ProjectName != nil {
		db = db.Where("Apply.project_name = ?", info.ProjectName)
	}
	if info.InsureName != nil {
		db = db.Where("Apply.insure_name = ?", info.InsureName)
	}
	//if info.ElogTemplateName != nil {
	//	db = db.Where("Apply.elog_template_name = ?", info.ElogTemplateName)
	//}
	if info.ElogNo != nil {
		db = db.Where("Letter.elog_no = ?", info.ElogNo)
	}
	//if info.OrderStatus != nil {
	//	db = db.Where("Apply.order_status = ?", info.OrderStatus)
	//}
	//if info.AuditStatus != nil {
	//	db = db.Where("Apply.audit_status = ?", info.AuditStatus)
	//}
	if info.OpenBeginDate != nil {
		db = db.Where("Apply.open_begin_date BETWEEN ? AND ?", info.OpenBeginDate[0], info.OpenBeginDate[1])
	}
	if info.ApplyCreatedAt != nil {
		db = db.Where("Apply.created_at BETWEEN ? AND ?", info.ApplyCreatedAt[0], info.ApplyCreatedAt[1])
	}
	if info.LetterCreatedAt != nil {
		db = db.Where("Letter.created_at BETWEEN ? AND ?", info.LetterCreatedAt[0], info.LetterCreatedAt[1])
	}
	if info.InsureDay != nil {
		db = db.Where("Letter.insure_day = ?", info.InsureDay)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Preload(clause.Associations).Order("order.created_at desc").Offset(offset).Find(&orders).Error
	return orders, total, err
}
