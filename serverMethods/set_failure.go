package serverMethods

import (
	"github.com/gin-gonic/gin"
	lib "gitlab.telemed.help/devops/ci/serverMethods/lib"
)

func SetFailure(c *gin.Context) {
	lib.SetPipelineSuccessStatus(c, false)
}
