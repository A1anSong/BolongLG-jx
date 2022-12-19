// 自动生成模板Template
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Template 结构体
type Template struct {
	global.GVA_MODEL
	TemplateName   *string `json:"templateName" form:"templateName" gorm:"type:varchar(128);"`
	TemplateFileID *uint   `json:"templateFileID" form:"templateFileID"`
	TemplateFile   *File   `json:"templateFile" form:"templateFile"`
}

// TableName Template 表名
func (Template) TableName() string {
	return "template"
}
