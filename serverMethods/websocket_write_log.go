package serverMethods

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	cfg "gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/models"
	lib "gitlab.telemed.help/devops/ci/serverMethods/lib"
	"gopkg.in/olahol/melody.v1"
)

var (
	logs lib.Logs
)

func init() {
	logs = lib.NewLogs()
}

func WSWriteLog(c *gin.Context) {
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

	if !pipeline.CompareToken(token) && c.ClientIP() != cfg.Get().LogWriterIPAddr {
		c.JSON(502, gin.H{
			"error": fmt.Sprintf("The token is not correct"),
		})
		return
	}

	wslogRouter := melody.New()

	wslogRouter.HandleConnect(func(s *melody.Session) {
		logs.Get(pipeline.Id)
	})

	wslogRouter.HandleDisconnect(func(s *melody.Session) {
		logs.Remove(pipeline.Id)
	})

	wslogRouter.HandleMessage(func(s *melody.Session, msg []byte) {
		logs.PushLine(pipeline.Id, string(msg))
	})

	wslogRouter.HandleRequest(c.Writer, c.Request)
}
