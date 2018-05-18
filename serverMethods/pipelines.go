package serverMethods

import (
	//"gitlab.telemed.help/devops/ci/serverMethods/helpers"
	"github.com/gin-gonic/gin"
)

func Pipelines(c *gin.Context) {
	//me := helpers.GetMe(c)

	c.JSON(200, gin.H{})
}
