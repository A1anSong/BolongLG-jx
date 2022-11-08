// 自动生成模板Revoke
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Revoke 结构体
type Revoke struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Revoke 表名
func (Revoke) TableName() string {
	return "revoke"
}
