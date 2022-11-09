// 自动生成模板Letter
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Letter 结构体
type Letter struct {
	global.GVA_MODEL
	OrderID             *uint    `json:"orderID" form:"orderID"`
	Order               *Order   `json:"-" form:"-"`
	OrderNo             *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ElogNo              *string  `json:"elogNo" form:"elogNo" gorm:"type:varchar(128);"`
	InsuranceName       *string  `json:"insuranceName" form:"insuranceName" gorm:"type:varchar(256);"`
	InsuranceCreditCode *string  `json:"insuranceCreditCode" form:"insuranceCreditCode" gorm:"type:varchar(18);"`
	EloOutDate          *string  `json:"eloOutDate" form:"eloOutDate" gorm:"type:varchar(19);"`
	EloUrl              *string  `json:"eloUrl" form:"eloUrl" gorm:"type:varchar(256);"`
	EloEncryptUrl       *string  `json:"eloEncryptUrl" form:"eloEncryptUrl" gorm:"type:varchar(256);"`
	TenderDeposit       *float64 `json:"tenderDeposit" form:"tenderDeposit"`
	InsureStartDate     *string  `json:"insureStartDate" form:"insureStartDate" gorm:"type:varchar(19);"`
	InsureEndDate       *string  `json:"insureEndDate" form:"insureEndDate" gorm:"type:varchar(19);"`
	InsureDay           *int64   `json:"insureDay" form:"insureDay"`
	ValidateCode        *string  `json:"validateCode" form:"validateCode" gorm:"type:varchar(8);"`
}

// TableName Letter 表名
func (Letter) TableName() string {
	return "letter"
}
