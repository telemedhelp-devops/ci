package serverMethods

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/xaionaro-go/extime"
	"github.com/xaionaro-go/log"
	"gitlab.telemed.help/devops/ci/models"
)

func WhatToDeploy(c *gin.Context) {
	var pipelines models.Pipelines
	var err error
	pipelines, err = models.PipelineSQL.Select("approved_at IS NOT NULL AND deleted_at IS NULL AND deploy_started_at IS NULL")
	if err != nil && err != sql.ErrNoRows {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	if pipelines == nil {
		pipelines = models.Pipelines{}
	}

	for _, pipeline := range pipelines {
		pipeline.DeployStartedAt = &[]extime.Time{extime.Now()}[0]
		err := pipeline.Update()
		if err != nil {
			log.Errorf("SQL error: %v", err)
		}
	}

	c.JSON(200, gin.H{
		"Pipelines": pipelines.HideTokenHash(),
	})
}
