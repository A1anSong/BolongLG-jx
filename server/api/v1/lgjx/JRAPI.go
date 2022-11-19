package lgjx

import "github.com/gin-gonic/gin"

type JRAPI struct {
}

// TODO:这里添加service
// var applyService = service.ServiceGroupApp.LgjxServiceGroup.ApplyService

func (jrApi *JRAPI) Apply(c *gin.Context) {
	c.String(200, "ok")
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
