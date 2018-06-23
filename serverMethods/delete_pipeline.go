package serverMethods

import (
	"github.com/gin-gonic/gin"
)

func DeletePipeline(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "pong",
	})
}
