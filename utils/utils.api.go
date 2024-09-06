package utils

import (
	"test-pari/schemas"

	"github.com/gin-gonic/gin"
)

func APIResponse(c *gin.Context, Code int64, Status string, Message string, Data interface{}) {
	c.JSON(int(Code), schemas.Response{Status: Status, Message: Message, Data: Data})
}
