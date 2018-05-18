package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	m "gitlab.telemed.help/devops/ci/models"
)

func GetMe(c *gin.Context) m.User {
	session := sessions.Default(c)
	return session.Get("User").(m.User)
}
