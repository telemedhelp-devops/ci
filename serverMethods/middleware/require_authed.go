package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	m "gitlab.telemed.help/devops/ci/models"
)

func init() {
}

func RequireAuthed(c *gin.Context) {
	session := sessions.Default(c)
	_, ok := session.Get("User").(m.User)
	if ok {
		c.Next()
		return
	}

	c.Abort()
	c.JSON(401, gin.H{
		"Error": "Authorization required",
	})
}
