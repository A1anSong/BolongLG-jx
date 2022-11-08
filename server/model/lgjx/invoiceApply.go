// 自动生成模板InvoiceApply
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// InvoiceApply 结构体
type InvoiceApply struct {
	global.GVA_MODEL
	ApplyNo string `json:"applyNo" form:"applyNo" gorm:"column:apply_no;comment:;"`
}

// TableName InvoiceApply 表名
func (InvoiceApply) TableName() string {
	return "invoice_apply"
}
