package serverMethods

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xaionaro-go/extime"
	"gitlab.telemed.help/devops/ci/models"
	"gitlab.telemed.help/devops/ci/serverMethods/helpers"
)

type PatchPipelineParams struct {
	Id int
}

func PatchPipeline(c *gin.Context) {
	me := helpers.GetMe(c)

	var params PatchPipelineParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Invalid JSON body: ` + err.Error(),
		})
		return
	}

	err = (&models.Approval{
		Username:   strings.ToLower(me.GitLabUser.NickName),
		PipelineId: params.Id,
		CreatedAt:  &[]extime.Time{extime.Now()}[0],
	}).Create()
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot create the approval: ` + err.Error(),
		})
		return
	}

	pipeline, err := models.PipelineSQL.First(models.Pipeline{Id: params.Id})
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot fetch the pipeline info: ` + err.Error(),
		})
		return
	}

	requiredApprovals, err := models.RequiredApprovalSQL.Select(models.RequiredApproval{ProjectName: pipeline.ProjectName})
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot fetch current required approvals: ` + err.Error(),
		})
		return
	}

	satisfiedApprovementGroupMap := map[int]bool{}
	usernameToApprovementGroupIdMap := map[string]int{}
	for _, requiredApproval := range requiredApprovals {
		satisfiedApprovementGroupMap[requiredApproval.ApprovementGroupId] = false
		usernameToApprovementGroupIdMap[strings.ToLower(requiredApproval.Username)] = requiredApproval.ApprovementGroupId
	}

	approvals, err := models.ApprovalSQL.Select(models.Approval{PipelineId: pipeline.Id})
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot fetch current approvals: ` + err.Error(),
		})
		return
	}

	for _, approval := range approvals {
		approvementGroupId := usernameToApprovementGroupIdMap[strings.ToLower(approval.Username)]
		satisfiedApprovementGroupMap[approvementGroupId] = true
	}

	for approvementGroupId, isApproved := range satisfiedApprovementGroupMap {
		if !isApproved {
			c.JSON(200, gin.H{
				"Status":  "ok",
				"Comment": fmt.Sprintf("approvement group %v has not approvals, yet", approvementGroupId),
			})
			return
		}
	}

	pipeline.ApprovedAt = &[]extime.Time{extime.Now()}[0]
	err = pipeline.Update()
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Cannot update the pipeline info: ` + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Status": "ok",
	})
}
