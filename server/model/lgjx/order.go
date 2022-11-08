// 自动生成模板Order
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Order 结构体
type Order struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "order"
}
