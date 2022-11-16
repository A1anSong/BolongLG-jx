package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	"time"
)

type OrderSearch struct {
	lgjx.Order
	ApplyNo          *string     `json:"applyNo" form:"applyNo"`
	ProjectName      *string     `json:"projectName" form:"projectName"`
	InsureName       *string     `json:"insureName" form:"insureName"`
	ElogTemplateName *string     `json:"elogTemplateName" form:"elogTemplateName"`
	ElogNo           *string     `json:"elogNo" form:"elogNo"`
	OrderStatus      *string     `json:"orderStatus" form:"orderStatus"`
	AuditStatus      *string     `json:"auditStatus" form:"auditStatus"`
	OpenBeginDate    []time.Time `json:"openBeginDate" form:"openBeginDate[]"`
	ApplyCreatedAt   []time.Time `json:"applyCreatedAt" form:"applyCreatedAt[]"`
	LetterCreatedAt  []time.Time `json:"letterCreatedAt" form:"letterCreatedAt[]"`
	InsureDay        *int        `json:"insureDay" form:"insureDay"`
	request.PageInfo
}
