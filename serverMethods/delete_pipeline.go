package serverMethods

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xaionaro-go/extime"
	"gitlab.telemed.help/devops/ci/models"
	"gitlab.telemed.help/devops/ci/serverMethods/helpers"
)

type DeletePipelineParams struct {
	Id int
}

func DeletePipeline(c *gin.Context) {
	me := helpers.GetMe(c)

	var params DeletePipelineParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Invalid JSON body: ` + err.Error(),
		})
		return
	}

	requiredApprovals, err := models.RequiredApprovalSQL.Select(models.RequiredApproval{Username: strings.ToLower(me.GitLabUser.NickName)})
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot get my projects: ` + err.Error(),
		})
		return
	}
	var rwProjectNames []string
	for _, requiredApproval := range requiredApprovals {
		rwProjectNames = append(rwProjectNames, requiredApproval.ProjectName)
	}

	pipeline, err := models.PipelineSQL.Where("`project_name` IN (?)", rwProjectNames).First(models.Pipeline{Id: params.Id})
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot delete the pipeline: ` + err.Error(),
		})
		return
	}

	pipeline.DeletedAt = &[]extime.Time{extime.Now()}[0]
	err = pipeline.Update()
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot delete the pipeline: ` + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "pong",
	})
}
