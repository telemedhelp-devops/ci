package main

import (
	"github.com/gin-gonic/gin"
	cfg "gitlab.telemed.help/devops/ci/config"
	db "gitlab.telemed.help/devops/ci/db"
	"gitlab.telemed.help/devops/ci/gitlabAuth"
)

func main() {
	cfg.Reload()
	db.Init()
	gitlabAuth.Init()

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
