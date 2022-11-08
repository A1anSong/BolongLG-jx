// 自动生成模板Apply
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Apply 结构体
type Apply struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Apply 表名
func (Apply) TableName() string {
	return "apply"
}
