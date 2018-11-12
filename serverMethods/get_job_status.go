package serverMethods

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/gitlab"
)

func GetJobStatus(c *gin.Context) {
	if c.ClientIP() != config.Get().LogReaderIPAddr {
		c.JSON(502, gin.H{
			"error": "Permission denied",
		})
		return
	}

	jobIdStr := c.Param("gitlab_job_id")
	jobId, err := strconv.Atoi(jobIdStr)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	projectIdStr := c.Param("gitlab_project_id")
	projectId, err := strconv.Atoi(projectIdStr)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	job, err := gitlab.GetJob(projectId, jobId)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"job": job,
	})
}
