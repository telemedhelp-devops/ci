package serverMethods

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"github.com/xaionaro-go/extime"
	"gitlab.telemed.help/devops/ci/models"
)

func WantToDeploy(c *gin.Context) {
	projectName := c.Param("project")
	tagName := c.Param("tag")

	pipeline, err := models.PipelineSQL.First(models.Pipeline{ProjectName: projectName, TagName: tagName})
	switch err {
	case nil:
	case sql.ErrNoRows:
		pipelineIdStr := c.Query("GitlabPipelineId")
		namespace := c.Query("GitlabNamespace")
		pipelineId, err := strconv.Atoi(pipelineIdStr)
		if err != nil {
			c.JSON(502, gin.H{
				"error": err.Error(),
			})
			return
		}
		token := c.Query("token")
		if token == "" {
			token, err = password.Generate(32, 10, 0, false, false)
			if err != nil {
				panic(err)
			}
		}
		pipeline := models.Pipeline{GitlabPipelineId: pipelineId, GitlabNamespace: namespace, ProjectName: projectName, TagName: tagName, CreatedAt: &[]extime.Time{extime.Now()}[0]}
		pipeline.SetToken(token)
		err = pipeline.Create()
		if err != nil {
			c.JSON(502, gin.H{
				"error": err.Error(),
			})
			return
		}
	default:
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Pipeline": pipeline.HideTokenHash(),
	})
}
