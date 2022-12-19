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

type TestJRAPI struct {
}

var testJRAPIService = service.ServiceGroupApp.LgjxTestServiceGroup.TestJRAPIService

func (testJRAPI *TestJRAPI) Apply(c *gin.Context) {
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
	if resApply, err = testJRAPIService.ApplyOrder(reApply); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
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

func (testJRAPI *TestJRAPI) PayPush(c *gin.Context) {
	var rePayPush jrrequest.JRAPIPayPush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &rePayPush)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resPayPush jrresponse.JRAPIPayPush
	if resPayPush, err = testJRAPIService.PayPush(rePayPush); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resPayPush); err != nil {
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

func (testJRAPI *TestJRAPI) QueryInfo(c *gin.Context) {
	var reQueryInfo jrrequest.JRAPIQueryInfo
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reQueryInfo)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resQueryInfo jrresponse.JRAPIQueryInfo
	if resQueryInfo, err = testJRAPIService.QueryInfo(reQueryInfo); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resQueryInfo); err != nil {
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

func (testJRAPI *TestJRAPI) RevokePush(c *gin.Context) {
	var reRevokePush jrrequest.JRAPIRevokePush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reRevokePush)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resRevokePush jrresponse.JRAPIRevokePush
	if resRevokePush, err = testJRAPIService.RevokePush(reRevokePush); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resRevokePush); err != nil {
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

func (testJRAPI *TestJRAPI) ApplyDelay(c *gin.Context) {
	var reApplyDelay jrrequest.JRAPIApplyDelay
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyDelay)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyDelay jrresponse.JRAPIApplyDelay
	if resApplyDelay, err = testJRAPIService.ApplyDelay(reApplyDelay); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resApplyDelay); err != nil {
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

func (testJRAPI *TestJRAPI) ApplyRefund(c *gin.Context) {
	var reApplyRefund jrrequest.JRAPIApplyRefund
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyRefund)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyRefund jrresponse.JRAPIApplyRefund
	if resApplyRefund, err = testJRAPIService.ApplyRefund(reApplyRefund); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resApplyRefund); err != nil {
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

func (testJRAPI *TestJRAPI) ApplyClaim(c *gin.Context) {
	var reApplyClaim jrrequest.JRAPIApplyClaim
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyClaim)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyClaim jrresponse.JRAPIApplyClaim
	if resApplyClaim, err = testJRAPIService.ApplyClaim(reApplyClaim); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resApplyClaim); err != nil {
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

func (testJRAPI *TestJRAPI) LogoutPush(c *gin.Context) {
	var reLogoutPush jrrequest.JRAPILogoutPush
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reLogoutPush)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resLogoutPush jrresponse.JRAPILogoutPush
	if resLogoutPush, err = testJRAPIService.LogoutPush(reLogoutPush); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resLogoutPush); err != nil {
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

func (testJRAPI *TestJRAPI) ApplyInvoice(c *gin.Context) {
	var reApplyInvoice jrrequest.JRAPIApplyInvoice
	jsonData, _ := c.Get("jsonData")
	err := json.Unmarshal([]byte(jsonData.(string)), &reApplyInvoice)
	if err != nil {
		global.GVA_LOG.Error("接收请求失败!", zap.Error(err))
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.FAILED),
			Msg:  jrapi.FAILED.String(),
			Data: err.Error(),
		})
		return
	}
	var resApplyInvoice jrresponse.JRAPIApplyInvoice
	if resApplyInvoice, err = testJRAPIService.ApplyInvoice(reApplyInvoice); err != nil {
		c.JSON(http.StatusOK, jrresponse.JRResponse{
			Code: int(jrapi.MissingServiceParam),
			Msg:  jrapi.MissingServiceParam.String(),
			Data: "",
		})
	} else {
		var response jrresponse.JRResponse
		if response, err = lgjx.GenJRResponse(resApplyInvoice); err != nil {
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

	if refile, err := testJRAPIService.DelayFileDownload(elog, encryptFlag); err != nil {
		c.String(http.StatusOK, "获取文件失败")
		return
	} else {
		c.Header("Content-Disposition", "attachment; filename="+*refile.FileName) // 用来指定下载下来的文件名
		c.Data(http.StatusOK, "application/octet-stream", refile.FileSteam)
	}
}

func (testJRAPI *TestJRAPI) InvoiceFileDownload(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
