package serverMethods

import (
	//"gitlab.telemed.help/devops/ci/serverMethods/helpers"
	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/models"
)

func Pipelines(c *gin.Context) {
	pipelines, err := models.PipelineSQL.Select()
	if err != nil {
		panic(err)
	}
	//me := helpers.GetMe(c)

	c.JSON(200, gin.H{
		"Pipelines": pipelines,
	})
}
