package serverMethods

import (
	"fmt"
	"strconv"

	cfg "gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func WSGetLog(c *gin.Context) {
	pipelineIdStr := c.Param("gitlab_pipeline_id")
	token := c.Query("token")

	if pipelineIdStr == "" {
		c.JSON(502, gin.H{
			"error": `pipelineIdStr == ""`,
		})
		return
	}

	pipelineId, err := strconv.Atoi(pipelineIdStr)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	pipelines, err := models.PipelineSQL.Select(models.Pipeline{GitlabPipelineId: pipelineId})
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(pipelines) != 1 {
		c.JSON(502, gin.H{
			"error": fmt.Sprintf("Found unexpected number of pipelines (should be only one): %v", len(pipelines)),
		})
		return
	}

	pipeline := pipelines[0]

	if !pipeline.CompareToken(token) && c.ClientIP() != cfg.Get().LogReaderIPAddr {
		c.JSON(502, gin.H{
			"error": fmt.Sprintf("The token is not correct"),
		})
		return
	}

	log := logs.Get(pipeline.Id)

	wslogRouter := melody.New()

	wslogRouter.HandleConnect(func(s *melody.Session) {
		log.AddSession(s)
	})

	wslogRouter.HandleDisconnect(func(s *melody.Session) {
		log.RemoveSession(s)
	})

	wslogRouter.HandleMessage(func(s *melody.Session, msg []byte) {
	})

	wslogRouter.HandleRequest(c.Writer, c.Request)
}
