package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	cfg "gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/gitlabAuth"
	m "gitlab.telemed.help/devops/ci/serverMethods"
	mw "gitlab.telemed.help/devops/ci/serverMethods/middleware"
)

func setupRouter(r *gin.Engine) {
	setupJsonRouter(r)
	setupFrontendRouter(r)
}

func setupJsonRouter(r *gin.Engine) {
	store := cookie.NewStore([]byte(cfg.Get().Secret))
	store.Options(sessions.Options{
		Secure:   !cfg.Get().IsDev,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400 * 30 * 3, // 3 months
	})

	gothic.Store = store

	r.Use(sessions.Sessions("session", store))

	authed := r.Group("/")
	authed.Use(mw.RequireAuthed) // some routines for an already authed

	// Auth
	r.GET("/auth/gitlab/login", gitlabAuth.Login)
	r.GET("/auth/gitlab/callback", gitlabAuth.Callback)

	// My methods
	r.GET("/ping.json", m.Ping)
	r.GET("/simpleApi/wantToDeploy/:project/:tag", m.WantToDeploy)
	authed.GET("/whoami.json", m.Whoami)
	authed.GET("/pipelines.json", m.Pipelines)
}

func setupFrontendRouter(r *gin.Engine) {
	r.Static("/frontend", "frontend/build")
	r.Static("/static", "frontend/build/static")
	r.Static("/css", "frontend/build/css")
	r.StaticFile("/", "frontend/build/index.html")
	//r.StaticFile("/projects/:project_id/approveDeploy", "frontend/build/index.html")
}
