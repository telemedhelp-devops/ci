package main

import (
	"github.com/gin-gonic/gin"
	cfg "gitlab.telemed.help/devops/ci/config"
	db "gitlab.telemed.help/devops/ci/db"
	"gitlab.telemed.help/devops/ci/gitlabAuth"
	models "gitlab.telemed.help/devops/ci/models"
	"gitlab.telemed.help/devops/ci/sms"
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

	models.ApproveTokenSQL.SetDefaultDB(db.GetDB())
	_, err = models.ApproveTokenSQL.Table().CreateTableIfNotExists(db.GetDB())
	checkError(err)

	models.UserProfileSQL.SetDefaultDB(db.GetDB())
	_, err = models.UserProfileSQL.Table().CreateTableIfNotExists(db.GetDB())
	checkError(err)

	err = sms.InitGoSMSC(cfg.Get().SMSGWLogin, cfg.Get().SMSGWPassword)
	checkError(err)

	gitlabAuth.Init()

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
