package middleware

import (
	"encoding/base64"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/lgjx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JRValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request jrrequest.JRRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingParam),
				Msg:  jrapi.MissingParam.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// sm3验签
		if request.Signature != lgjx.SM3Sign("appKey="+request.AppKey+"&data="+request.Data+"&requestId="+request.RequestId+"&timestamp="+request.TimeStamp) {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.SignCheckFailed),
				Msg:  jrapi.SignCheckFailed.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// 提取data
		byteEncryptData, err := base64.StdEncoding.DecodeString(request.Data)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.MissingParam),
				Msg:  jrapi.MissingParam.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		// sm4解密
		jsonData, err := lgjx.Sm4Decrypt(byteEncryptData)
		if err != nil {
			c.JSON(http.StatusOK, jrresponse.JRResponse{
				Code: int(jrapi.DecryptFailed),
				Msg:  jrapi.DecryptFailed.String(),
				Data: "",
			})
			c.Abort()
			return
		}
		c.Set("jsonData", jsonData)
		c.Next()
	}
}
