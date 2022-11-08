package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
)

type TestInvoiceApplyService struct {
}

func (testInvoiceApplyService *TestInvoiceApplyService) CreateInvoiceApply(invoiceApply lgjx.InvoiceApply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&invoiceApply).Error
	return err
}

func (testInvoiceApplyService *TestInvoiceApplyService) DeleteInvoiceApply(invoiceApply lgjx.InvoiceApply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&invoiceApply).Error
	return err
}

func (testInvoiceApplyService *TestInvoiceApplyService) DeleteInvoiceApplyByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.InvoiceApply{}, "id in ?", ids.Ids).Error
	return err
}

func (testInvoiceApplyService *TestInvoiceApplyService) UpdateInvoiceApply(invoiceApply lgjx.InvoiceApply) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&invoiceApply).Error
	return err
}

func (testInvoiceApplyService *TestInvoiceApplyService) GetInvoiceApply(id uint) (invoiceApply lgjx.InvoiceApply, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&invoiceApply).Error
	return
}

func (testInvoiceApplyService *TestInvoiceApplyService) GetInvoiceApplyInfoList(info lgjxReq.InvoiceApplySearch) (list []lgjx.InvoiceApply, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.InvoiceApply{})
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
