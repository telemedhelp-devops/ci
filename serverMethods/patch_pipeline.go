package serverMethods

import (
	"github.com/gin-gonic/gin"
)

func PatchPipeline(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "pong",
	})
}
