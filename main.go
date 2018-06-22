package main

import (
	"github.com/gin-gonic/gin"
	cfg "gitlab.telemed.help/devops/ci/config"
	db "gitlab.telemed.help/devops/ci/db"
	models "gitlab.telemed.help/devops/ci/models"
	"gitlab.telemed.help/devops/ci/gitlabAuth"
)

func main() {
	cfg.Reload()
	db.Init()
	models.PipelineSQL.SetDefaultDB(db.GetDB())
	_, err := models.PipelineSQL.Table().CreateTableIfNotExists(db.GetDB())
	if err != nil {
		panic(err)
	}
	gitlabAuth.Init()

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
