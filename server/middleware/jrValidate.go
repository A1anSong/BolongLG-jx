package middleware

import (
	"encoding/base64"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrresponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JRValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request jrrequest.JRRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
				Code: int(jrapi.MissingParam),
				Msg:  jrapi.MissingParam.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// TODO: 这里写sm3验证流程
		byteEncryptData, err := base64.StdEncoding.DecodeString(request.Data)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponseFailed{
				Code: int(jrapi.MissingParam),
				Msg:  jrapi.MissingParam.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// TODO: 这里写sm4解密流程
		jsonData := byteEncryptData
		c.Set("jsonData", jsonData)
		c.Next()
	}
}
