package serverMethods

import (
	"github.com/gin-gonic/gin"
	lib "gitlab.telemed.help/devops/ci/serverMethods/lib"
)

func SetSuccess(c *gin.Context) {
	lib.SetPipelineSuccessStatus(c, true)
}
