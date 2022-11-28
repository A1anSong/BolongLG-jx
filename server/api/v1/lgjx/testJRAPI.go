package lgjx

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestJRAPI struct {
}

var testJRAPIService = service.ServiceGroupApp.LgjxTestServiceGroup.TestJRAPIService

func (testJRAPI *TestJRAPI) Apply(c *gin.Context) {
	var reApply jrrequest.JRAPIApply
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reApply)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resApply jrresponse.JRAPIApply
	if resApply, err = testJRAPIService.ApplyOrder(reApply); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resApply)
	}
}

func (testJRAPI *TestJRAPI) PayPush(c *gin.Context) {
	var rePayPush jrrequest.JRAPIPayPush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &rePayPush)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resPayPush jrresponse.JRAPIPayPush
	if resPayPush, err = testJRAPIService.PayPush(rePayPush); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resPayPush)
	}
}

func (testJRAPI *TestJRAPI) QueryInfo(c *gin.Context) {
	var reQueryInfo jrrequest.JRAPIQueryInfo
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reQueryInfo)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resQueryInfo jrresponse.JRAPIQueryInfo
	if resQueryInfo, err = testJRAPIService.QueryInfo(reQueryInfo); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resQueryInfo)
	}
}

func (testJRAPI *TestJRAPI) RevokePush(c *gin.Context) {
	var reRevokePush jrrequest.JRAPIRevokePush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reRevokePush)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resRevokePush jrresponse.JRAPIRevokePush
	if resRevokePush, err = testJRAPIService.RevokePush(reRevokePush); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resRevokePush)
	}
}

func (testJRAPI *TestJRAPI) ApplyDelay(c *gin.Context) {
	var reApplyDelay jrrequest.JRAPIApplyDelay
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reApplyDelay)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resApplyDelay jrresponse.JRAPIApplyDelay
	if resApplyDelay, err = testJRAPIService.ApplyDelay(reApplyDelay); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resApplyDelay)
	}
}

func (testJRAPI *TestJRAPI) ApplyRefund(c *gin.Context) {
	var reApplyRefund jrrequest.JRAPIApplyRefund
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reApplyRefund)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resApplyRefund jrresponse.JRAPIApplyRefund
	if resApplyRefund, err = testJRAPIService.ApplyRefund(reApplyRefund); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resApplyRefund)
	}
}

func (testJRAPI *TestJRAPI) ApplyClaim(c *gin.Context) {
	var reApplyClaim jrrequest.JRAPIApplyClaim
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reApplyClaim)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resApplyClaim jrresponse.JRAPIApplyClaim
	if resApplyClaim, err = testJRAPIService.ApplyClaim(reApplyClaim); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resApplyClaim)
	}
}

func (testJRAPI *TestJRAPI) LogoutPush(c *gin.Context) {
	var reLogoutPush jrrequest.JRAPILogoutPush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reLogoutPush)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resLogoutPush jrresponse.JRAPILogoutPush
	if resLogoutPush, err = testJRAPIService.LogoutPush(reLogoutPush); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resLogoutPush)
	}
}

func (testJRAPI *TestJRAPI) ApplyInvoice(c *gin.Context) {
	var reApplyInvoice jrrequest.JRAPIApplyInvoice
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal(jsonData.([]byte), &reApplyInvoice)
	if err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
		return
	}
	var resApplyInvoice jrresponse.JRAPIApplyInvoice
	if resApplyInvoice, err = testJRAPIService.ApplyInvoice(reApplyInvoice); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		c.JSON(http.StatusOK, resApplyInvoice)
	}
}

func (testJRAPI *TestJRAPI) LetterFileDownload(c *gin.Context) {
	elog, ok := c.GetQuery("elog")
	if !ok {
		c.String(http.StatusOK, "非法参数")
		return
	}
	encrypt, ok := c.GetQuery("type")
	encryptFlag := false
	if ok {
		if encrypt == "encrypt" {
			encryptFlag = true
		} else {
			c.String(http.StatusOK, "非法参数")
			return
		}
	}

	if refile, err := testJRAPIService.LetterFileDownload(elog, encryptFlag); err != nil {
		c.String(http.StatusOK, "获取文件失败")
		return
	} else {
		c.Header("Content-Disposition", "attachment; filename="+*refile.FileName) // 用来指定下载下来的文件名
		c.Data(http.StatusOK, "application/octet-stream", refile.FileSteam)
	}
}

func (testJRAPI *TestJRAPI) DelayFileDownload(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

func (testJRAPI *TestJRAPI) InvoiceFileDownload(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
