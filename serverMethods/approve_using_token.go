package serverMethods

import (
	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/models"
)

func ApproveUsingToken(c *gin.Context) {
	tokenValue := c.Param("token")

	if tokenValue == "" {
		c.JSON(502, gin.H{
			"error": "token is empty",
		})
	}

	token, err := models.ApproveTokenSQL.First(models.ApproveToken{Token: tokenValue})
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.GetUserByUsername(token.Username)
	err = user.ApprovePipeline(token.PipelineId)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}

	token.Delete()

	c.JSON(200, gin.H{
		"Status": "ok",
	})
}
