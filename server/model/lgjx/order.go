// 自动生成模板Order
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Order 结构体
type Order struct {
	global.GVA_MODEL
	OrderNo   *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ApplyID   *uint    `json:"-" form:"-"`
	Apply     *Apply   `json:"apply" form:"apply"`
	PayID     *uint    `json:"-" form:"-"`
	Pay       *Pay     `json:"pay" form:"pay"`
	LetterID  *uint    `json:"-" form:"-"`
	Letter    *Letter  `json:"letter" form:"letter"`
	RevokeID  *uint    `json:"-" form:"-"`
	Revoke    *Revoke  `json:"revoke" form:"revoke"`
	DelayID   *uint    `json:"-" form:"-"`
	Delay     *Delay   `json:"delay" form:"delay"`
	RefundID  *uint    `json:"-" form:"-"`
	Refund    *Refund  `json:"refund" form:"refund"`
	LogoutID  *uint    `json:"-" form:"-"`
	Logout    *Logout  `json:"logout" form:"logout"`
	InvoiceID *uint    `json:"-" form:"-"`
	Invoice   *Invoice `json:"invoice" form:"invoice"`
	ProjectID *uint    `json:"-" form:"-"`
	Project   *Project `json:"project" form:"project"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "order"
}
