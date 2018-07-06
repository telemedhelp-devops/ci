package serverMethods

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/models"
)

func GetStatus(c *gin.Context) {
	pipelineIdStr := c.Param("gitlab_pipeline_id")
	pipelineId, err := strconv.Atoi(pipelineIdStr)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	pipeline, err := models.PipelineSQL.First(models.Pipeline{GitlabPipelineId: pipelineId})
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot fetch pipelines: ` + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": pipeline.Success,
	})
}
