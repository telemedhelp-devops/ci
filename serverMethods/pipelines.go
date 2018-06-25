package serverMethods

import (
	"database/sql"

	//"gitlab.telemed.help/devops/ci/serverMethods/helpers"
	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/models"
)

func Pipelines(c *gin.Context) {
	var err error
	var pipelines models.Pipelines
	pipelines, err = models.PipelineSQL.Select("`deleted_at` IS NULL AND `approved_at` IS NULL")
	if err != nil && err != sql.ErrNoRows {
		c.JSON(502, gin.H{
			"error": `Cannot fetch pipelines: `+err.Error(),
		})
		return
	}

	if pipelines == nil {
		pipelines = models.Pipelines{}
	}

	pipelines.PrepareApprovals()
	pipelines.PrepareRequiredApprovals()

	//me := helpers.GetMe(c)

	c.JSON(200, gin.H{
		"Pipelines": pipelines,
	})
}
