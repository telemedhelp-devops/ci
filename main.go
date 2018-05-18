package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	cfg "gitlab.telemed.help/devops/ci/config"
	db "gitlab.telemed.help/devops/ci/db"
)

func main() {
	cfg.Reload()
	db.Init()
	sessions.NewCookieStore([]byte(cfg.Get().Secret))

	r := gin.Default()
	setupRouter(r)
	r.Run()
}
