package serverMethodsLib

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xaionaro-go/extime"
	"gitlab.telemed.help/devops/ci/models"
)

func SetPipelineSuccessStatus(c *gin.Context, v bool) {
	pipelineIdStr := c.Param("gitlab_pipeline_id")
	pipelineId, err := strconv.Atoi(pipelineIdStr)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	pipeline, err := models.PipelineSQL.First(models.Pipeline{GitlabPipelineId: pipelineId})
	if err != nil && err != sql.ErrNoRows {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	pipeline.DeletedAt = &[]extime.Time{extime.Now()}[0]
	pipeline.Success = &v
	err = pipeline.Update()
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "OK",
	})
}
