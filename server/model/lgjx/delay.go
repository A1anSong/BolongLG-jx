// 自动生成模板Delay
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Delay 结构体
type Delay struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Delay 表名
func (Delay) TableName() string {
	return "delay"
}
