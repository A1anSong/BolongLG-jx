// 自动生成模板Logout
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Logout 结构体
type Logout struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Logout 表名
func (Logout) TableName() string {
	return "logout"
}
