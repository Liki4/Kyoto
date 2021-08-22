package toolkit

import (
	"github.com/gin-gonic/gin"
)

func Entry(handler func(ctx *gin.Context) (int, int, interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		statusCode, code, data := handler(c)
		if statusCode != 200 {
			c.JSON(statusCode, MakeErrorReturn(code, data))
		} else {
			c.JSON(statusCode, MakeSuccessReturn(data))
		}
	}
}
