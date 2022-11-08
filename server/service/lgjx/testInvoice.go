package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestInvoiceService struct {
}

func (testInvoiceService *TestInvoiceService) CreateInvoice(invoice lgjx.Invoice) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&invoice).Error
	return err
}

func (testInvoiceService *TestInvoiceService) DeleteInvoice(invoice lgjx.Invoice) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&invoice).Error
	return err
}

func (testInvoiceService *TestInvoiceService) DeleteInvoiceByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Invoice{}, "id in ?", ids.Ids).Error
	return err
}

func (testInvoiceService *TestInvoiceService) UpdateInvoice(invoice lgjx.Invoice) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&invoice).Error
	return err
}

func (testInvoiceService *TestInvoiceService) GetInvoice(id uint) (invoice lgjx.Invoice, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&invoice).Error
	return
}

func (testInvoiceService *TestInvoiceService) GetInvoiceInfoList(info lgjxReq.InvoiceSearch) (list []lgjx.Invoice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Invoice{})
	var invoices []lgjx.Invoice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&invoices).Error
	return invoices, total, err
}
