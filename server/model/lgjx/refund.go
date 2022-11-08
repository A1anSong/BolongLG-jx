// 自动生成模板Refund
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Refund 结构体
type Refund struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Refund 表名
func (Refund) TableName() string {
	return "refund"
}
