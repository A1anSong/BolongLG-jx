// 自动生成模板Pay
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Pay 结构体
type Pay struct {
	global.GVA_MODEL
	OrderID    *uint    `json:"orderID" form:"orderID"`
	Order      *Order   `json:"-" form:"-"`
	OrderNo    *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	PayNo      *string  `json:"payNo" form:"payNo" gorm:"type:varchar(64);"`
	PayAmount  *float64 `json:"payAmount" form:"payAmount"`
	PayTime    *string  `json:"payTime" form:"payTime" gorm:"type:varchar(19);"`
	PayTransNo *string  `json:"payTransNo" form:"payTransNo" gorm:"type:varchar(256);"`
}

// TableName Pay 表名
func (Pay) TableName() string {
	return "pay"
}
