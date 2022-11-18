package lgjx

import "github.com/gin-gonic/gin"

type TestJRAPI struct {
}

// TODO 这里添加service
// var applyService = service.ServiceGroupApp.LgjxServiceGroup.ApplyService

func (testJRAPI *TestJRAPI) Apply(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) PayPush(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) QueryInfo(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) RevokePush(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) ApplyDelay(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) ApplyRefund(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) ApplyClaim(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) LogoutPush(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) ApplyInvoice(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) LetterFileDownload(c *gin.Context) {
	c.String(200, "ok")
}

func (testJRAPI *TestJRAPI) DelayFileDownload(c *gin.Context) {
	c.String(200, "ok")
}
