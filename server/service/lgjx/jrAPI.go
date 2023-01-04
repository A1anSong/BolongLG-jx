package lgjx

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrresponse"
)

type JRAPIService struct {
}

func (jrAPIService *JRAPIService) ApplyOrder(reApply jrrequest.JRAPIApply) (resApply jrresponse.JRAPIApply, err error) {
	return resApply, errors.New("待添加")
}
