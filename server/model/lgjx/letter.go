// 自动生成模板Letter
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Letter 结构体
type Letter struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Letter 表名
func (Letter) TableName() string {
	return "letter"
}
