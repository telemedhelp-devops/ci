package serverMethods

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/models"
)

func WhatToDeploy(c *gin.Context) {
	var pipelines models.Pipelines
	var err error
	pipelines, err = models.PipelineSQL.Select("approved_at IS NOT NULL AND deleted_at IS NULL")
	if err != nil && err != sql.ErrNoRows {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	if pipelines == nil {
		pipelines = models.Pipelines{}
	}

	c.JSON(200, gin.H{
		"Pipelines": pipelines.HideTokenHash(),
	})
}
