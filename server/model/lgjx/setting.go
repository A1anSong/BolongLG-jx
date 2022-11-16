// 自动生成模板Setting
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Setting 结构体
type Setting struct {
	global.GVA_MODEL
	CompanyName string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:;"`
}

// TableName Setting 表名
func (Setting) TableName() string {
	return "setting"
}
