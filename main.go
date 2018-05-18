package main

import (
	"github.com/gin-gonic/gin"
	cfg "gitlab.telemed.help/devops/ci/config"
	db "gitlab.telemed.help/devops/ci/db"
)

func main() {
	cfg.Reload()
	db.Init()

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
