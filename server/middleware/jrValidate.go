package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func JRValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 添加验签规则
		fmt.Println("模拟验签")
	}
}
