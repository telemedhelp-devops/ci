package serverMethodsLib

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xaionaro-go/extime"
	logger "github.com/xaionaro-go/log"
	"gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/gitlab"
	"gitlab.telemed.help/devops/ci/models"
	"gitlab.telemed.help/devops/ci/slack"
	"gitlab.telemed.help/devops/ci/smtp"
)

func SetPipelineSuccessStatus(c *gin.Context, v bool) {
	cfg := config.Get()

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

	var slackErr error
	if v == true {
		tagDescription := ""
		tag, err := gitlab.GetTag(pipeline.GitlabNamespace+"/"+pipeline.ProjectName, pipeline.TagName)
		if err != nil {
			logger.Errorf("Cannot get the tag info from the gitlab: %v", err)
		} else {
			if tag.Release == nil {
				tagDescription = tag.Release.Description
			} else {
				logger.Errorf("Cannot get the release description")
			}
		}

		slackErr = slack.Send("Deployed (pipeline " + pipeline.InfoMarkdown() + ")\n\n" + tagDescription)
	} else {
		slackErr = slack.Send("Failed to deploy (pipeline " + pipeline.InfoMarkdown() + ")")
	}
	if slackErr != nil {
		logger.Errorf("Cannot send to the slack/mattermost: %v", slackErr)
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

	tag, err := gitlab.GetTag(pipeline.GitlabNamespace+"/"+pipeline.ProjectName, pipeline.TagName)
	if err != nil {
		logger.Errorf("Cannot find the tag on the GitLab: %v", err)
	}

	tagDescription := ""
	if tag.Release == nil {
		logger.Errorf("tag.Release == nil")
	} else {
		tagDescription = tag.Release.Description
	}

	notificationEmail := cfg.NotificationEmail

	title := fmt.Sprintf(`"%v/%v" has been deployed`, pipeline.ProjectName, pipeline.TagName)
	message := title + ".\n\n" + tagDescription
	smtp.Send(notificationEmail, title, message)

	c.JSON(200, gin.H{
		"status": "OK",
	})
}
