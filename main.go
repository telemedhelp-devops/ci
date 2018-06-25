package main

import (
	"github.com/gin-gonic/gin"
	cfg "gitlab.telemed.help/devops/ci/config"
	db "gitlab.telemed.help/devops/ci/db"
	"gitlab.telemed.help/devops/ci/gitlabAuth"
	models "gitlab.telemed.help/devops/ci/models"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	cfg.Reload()
	db.Init()

	models.PipelineSQL.SetDefaultDB(db.GetDB())
	_, err = models.PipelineSQL.Table().CreateTableIfNotExists(db.GetDB())
	checkError(err)

	models.ApprovalSQL.SetDefaultDB(db.GetDB())
	_, err = models.ApprovalSQL.Table().CreateTableIfNotExists(db.GetDB())
	checkError(err)

	models.RequiredApprovalSQL.SetDefaultDB(db.GetDB())
	_, err = models.RequiredApprovalSQL.Table().CreateTableIfNotExists(db.GetDB())
	checkError(err)

	gitlabAuth.Init()

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
