package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type InvoiceApplyService struct {
}

// CreateInvoiceApply 创建InvoiceApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceApplyService *InvoiceApplyService) CreateInvoiceApply(invoiceApply lgjx.InvoiceApply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Create(&invoiceApply).Error
	return err
}

// DeleteInvoiceApply 删除InvoiceApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceApplyService *InvoiceApplyService) DeleteInvoiceApply(invoiceApply lgjx.InvoiceApply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&invoiceApply).Error
	return err
}

// DeleteInvoiceApplyByIds 批量删除InvoiceApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceApplyService *InvoiceApplyService) DeleteInvoiceApplyByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Delete(&[]lgjx.InvoiceApply{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateInvoiceApply 更新InvoiceApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceApplyService *InvoiceApplyService) UpdateInvoiceApply(invoiceApply lgjx.InvoiceApply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Save(&invoiceApply).Error
	return err
}

// GetInvoiceApply 根据id获取InvoiceApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceApplyService *InvoiceApplyService) GetInvoiceApply(id uint) (invoiceApply lgjx.InvoiceApply, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx").Where("id = ?", id).First(&invoiceApply).Error
	return
}

// GetInvoiceApplyInfoList 分页获取InvoiceApply记录
// Author [piexlmax](https://github.com/piexlmax)
func (invoiceApplyService *InvoiceApplyService) GetInvoiceApplyInfoList(info lgjxReq.InvoiceApplySearch) (list []lgjx.InvoiceApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx").Model(&lgjx.InvoiceApply{})
	var invoiceApplys []lgjx.InvoiceApply
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&invoiceApplys).Error
	return invoiceApplys, total, err
}
