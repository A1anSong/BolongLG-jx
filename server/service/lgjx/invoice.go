package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type InvoiceService struct {
}

// CreateInvoice 创建Invoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceService *InvoiceService) CreateInvoice(invoice lgjx.Invoice) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&invoice).Error
	return err
}

// DeleteInvoice 删除Invoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceService *InvoiceService) DeleteInvoice(invoice lgjx.Invoice) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&invoice).Error
	return err
}

// DeleteInvoiceByIds 批量删除Invoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceService *InvoiceService) DeleteInvoiceByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.Invoice{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateInvoice 更新Invoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceService *InvoiceService) UpdateInvoice(invoice lgjx.Invoice) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&invoice).Error
	return err
}

// GetInvoice 根据id获取Invoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceService *InvoiceService) GetInvoice(id uint) (invoice lgjx.Invoice, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&invoice).Error
	return
}

// GetInvoiceInfoList 分页获取Invoice记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceService *InvoiceService) GetInvoiceInfoList(info lgjxReq.InvoiceSearch) (list []lgjx.Invoice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.Invoice{})
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
