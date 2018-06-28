package serverMethods

import (
	"github.com/gin-gonic/gin"
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

	err = me.ApprovePipeline(params.Id)
	if err != nil {
		c.JSON(502, gin.H{
			"error": `Got error while approvement: ` + err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Status": "ok",
	})
}
