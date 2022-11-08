// 自动生成模板Pay
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Pay 结构体
type Pay struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Pay 表名
func (Pay) TableName() string {
	return "pay"
}
