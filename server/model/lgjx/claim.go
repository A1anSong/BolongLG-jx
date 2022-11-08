// 自动生成模板Claim
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Claim 结构体
type Claim struct {
	global.GVA_MODEL
	OrderNo string `json:"orderNo" form:"orderNo" gorm:"column:order_no;comment:;"`
}

// TableName Claim 表名
func (Claim) TableName() string {
	return "claim"
}
