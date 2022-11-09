// 自动生成模板Apply
package lgjx

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Apply 结构体
type Apply struct {
	global.GVA_MODEL
	OrderID             *uint    `json:"orderID" form:"orderID"`
	Order               *Order   `json:"-" form:"-"`
	OrderNo             *string  `json:"orderNo" form:"orderNo" gorm:"type:varchar(64);"`
	ApplyNo             *string  `json:"applyNo" form:"applyNo" gorm:"type:varchar(64);"`
	ProductNo           *string  `json:"productNo" form:"productNo" gorm:"type:varchar(64);"`
	ProductType         *int64   `json:"productType" form:"productType"`
	ProductRate         *float64 `json:"productRate" form:"productRate"`
	ElogAmount          *float64 `json:"elogAmount" form:"elogAmount"`
	ProjectGuid         *string  `json:"projectGuid" form:"projectGuid" gorm:"type:varchar(64);"`
	ProjectName         *string  `json:"projectName" form:"projectName" gorm:"type:varchar(256);"`
	ProjectNo           *string  `json:"projectNo" form:"projectNo" gorm:"varchar(128);"`
	TenderDeposit       *float64 `json:"tenderDeposit" form:"tenderDeposit"`
	DepositStartDate    *string  `json:"depositStartDate" form:"depositStartDate" gorm:"varchar(19);"`
	DepositEndDate      *string  `json:"depositEndDate" form:"depositEndDate" gorm:"varchar(19);"`
	OpenBeginDate       *string  `json:"openBeginDate" form:"openBeginDate" gorm:"varchar(19);"`
	ElogTemplateNo      *string  `json:"elogTemplateNo" form:"elogTemplateNo" gorm:"varchar(14);"`
	ElogTemplateName    *string  `json:"elogTemplateName" form:"elogTemplateName" gorm:"varchar(256);"`
	InsuredName         *string  `json:"insuredName" form:"insuredName" gorm:"varchar(256);"`
	InsuredCreditCode   *string  `json:"insuredCreditCode" form:"insuredCreditCode" gorm:"varchar(18);"`
	InsuredAddress      *string  `json:"insuredAddress" form:"insuredAddress" gorm:"varchar(256);"`
	InsureName          *string  `json:"insureName" form:"insureName" gorm:"varchar(256);"`
	InsureCreditCode    *string  `json:"insureCreditCode" form:"insureCreditCode" gorm:"varchar(18);"`
	InsureLegalName     *string  `json:"insureLegalName" form:"insureLegalName" gorm:"varchar(64);"`
	InsureLegalIdCard   *string  `json:"insureLegalIdCard" form:"insureLegalIdCard" gorm:"varchar(18);"`
	InsureAddress       *string  `json:"insureAddress" form:"insureAddress" gorm:"varchar(256);"`
	ApplicantName       *string  `json:"applicantName" form:"applicantName" gorm:"varchar(64);"`
	ApplicantIdCard     *string  `json:"applicantIdCard" form:"applicantIdCard" gorm:"varchar(18);"`
	ApplicantTel        *string  `json:"applicantTel" form:"applicantTel" gorm:"varchar(11);"`
	AttachInfo          *string  `json:"attachInfo" form:"attachInfo" gorm:"text;"`
	AuditStatus         *int64   `json:"auditStatus" form:"auditStatus"`
	AuditOpinion        *string  `json:"auditOpinion" form:"auditOpinion" gorm:"varchar(512);"`
	AuditDate           *string  `json:"auditDate" form:"auditDate" gorm:"varchar(19);"`
	RealElogAmount      *float64 `json:"realElogAmount" form:"realElogAmount"`
	RealElogRate        *float64 `json:"realElogRate" form:"realElogRate"`
	InsuranceName       *string  `json:"insuranceName" form:"insuranceName" gorm:"varchar(256);"`
	InsuranceCreditCode *string  `json:"insuranceCreditCode" form:"insuranceCreditCode" gorm:"varchar(18);"`
}

// TableName Apply 表名
func (Apply) TableName() string {
	return "apply"
}
