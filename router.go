package main

import (
	m "gitlab.telemed.help/devops/ci/serverMethods"
	"gitlab.telemed.help/devops/ci/gitlabAuth"
	mw "gitlab.telemed.help/devops/ci/serverMethods/middleware"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine) {
	setupJsonRouter(r)
	setupFrontendRouter(r)
}

func setupJsonRouter(r *gin.Engine) {
	authed := r.Group("/")
	authed.Use(mw.RequireAuthed) // some routines for an already authed

	// Auth
	r.GET("/gitlab/callback", gitlabAuth.Callback)

	// My methods
	r.GET("/ping.json", m.Ping)
	authed.GET("/whoami.json", m.Whoami)
}

func setupFrontendRouter(r *gin.Engine) {
	r.Static("/frontend", "frontend/build")
	r.Static("/static", "frontend/build/static")
	r.Static("/css", "frontend/build/css")
	r.StaticFile("/", "frontend/build/index.html")
	//r.StaticFile("/projects/:project_id/approveDeploy", "frontend/build/index.html")
}
