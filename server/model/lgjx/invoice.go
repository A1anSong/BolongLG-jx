// 自动生成模板Invoice
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Invoice 结构体
type Invoice struct {
	global.GVA_MODEL
	InvoiceNo string `json:"invoiceNo" form:"invoiceNo" gorm:"column:invoice_no;comment:;"`
}

// TableName Invoice 表名
func (Invoice) TableName() string {
	return "invoice"
}
