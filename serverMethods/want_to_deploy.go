package serverMethods

import (
	"database/sql"
	"github.com/gin-gonic/gin"
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
		pipeline := models.Pipeline{ProjectName: projectName, TagName: tagName, CreatedAt: &[]extime.Time{extime.Now()}[0]}
		err := pipeline.Create()
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
		"Pipeline": pipeline,
	})
}
