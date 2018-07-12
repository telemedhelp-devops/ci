package serverMethods

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xaionaro-go/log"
	"gitlab.telemed.help/devops/ci/models"
	"gitlab.telemed.help/devops/ci/serverMethods/helpers"
)

func Approve(c *gin.Context) {
	me := helpers.GetMe(c)

	pipelineIdStr := c.Param("pipeline_id")
	pipelineId, err := strconv.Atoi(pipelineIdStr)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = me.ApprovePipeline(pipelineId)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	tokens, err := models.ApproveTokenSQL.Select(models.ApproveToken{Username: me.GetUsername(), PipelineId: pipelineId})
	if err != nil {
		log.Errorf("Cannot delete stale tokens: %v", err)
	} else {
		for _, token := range tokens {
			token.Delete()
		}
	}

	c.JSON(200, gin.H{
		"Status": "ok",
	})
}
