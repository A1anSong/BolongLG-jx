package lgjx

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/lgjx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type JRAPI struct {
}

var JRAPIService = service.ServiceGroupApp.LgjxServiceGroup.JRAPIService

func (jrApi *JRAPI) Apply(c *gin.Context) {
	var reApply jrrequest.JRAPIApply
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApply)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApply jrresponse.JRAPIApply
	if resApply, err = JRAPIService.ApplyOrder(reApply); err != nil {
		global.GVA_LOG.Error("创建Apply失败!", zap.Error(err))
		if err.Error() == jrapi.MissingServiceParam.String() {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingServiceParam),
				Msg:  jrapi.MissingServiceParam.String(),
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.FAILED),
				Msg:  err.Error(),
				Data: "",
			})
		}
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resApply); err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: -1,
				Msg:  "resApply序列化失败",
				Data: "",
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func (jrApi *JRAPI) PayPush(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) QueryInfo(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) RevokePush(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) ApplyDelay(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) ApplyRefund(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) ApplyClaim(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) LogoutPush(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) ApplyInvoice(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) LetterFileDownload(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) DelayFileDownload(c *gin.Context) {
	c.String(200, "ok")
}

func (jrApi *JRAPI) InvoiceFileDownload(c *gin.Context) {
	c.String(200, "ok")
}
