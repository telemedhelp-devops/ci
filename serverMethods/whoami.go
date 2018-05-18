package serverMethods

import (
	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/serverMethods/helpers"
)

func Whoami(c *gin.Context) {
	me := helpers.GetMe(c)

	c.JSON(200, gin.H{
		"User": me,
	})
}
