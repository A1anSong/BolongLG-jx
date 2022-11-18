// 自动生成模板File
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// File 结构体
type File struct {
	global.GVA_MODEL
	FileName  *string `json:"fileName" form:"fileName" gorm:"type:varchar(128);"`
	FileSteam []byte  `json:"-" form:"-" gorm:"type:mediumblob"`
}

// TableName File 表名
func (File) TableName() string {
	return "file"
}
