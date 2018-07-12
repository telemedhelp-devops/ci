package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gitlab.telemed.help/devops/ci/gitlabAuth"
	m "gitlab.telemed.help/devops/ci/models"
	"net/http"
)

var (
	ErrAuthRequired = fmt.Errorf("Authorization required")
)

func trySetUser(c *gin.Context) error {
	session := sessions.Default(c)
	userJson, ok := session.Get("User").(string)
	if !ok {
		return ErrAuthRequired
	}

	var user m.User
	err := json.Unmarshal(([]byte)(userJson), &user)
	if err != nil {
		return fmt.Errorf("Authorization required (cannot unserialize user data): %v", err.Error())
	}

	c.Set("User", user)
	return nil
}

func RequireAuthed(c *gin.Context) {
	err := trySetUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": err.Error(),
		})
		c.Abort()
		return
	}

	c.Next()
}

func RequireAuthedOrRedirect(c *gin.Context) {
	err := trySetUser(c)
	if err != nil {
		// redirect
		gitlabAuth.Login(c)
		c.Abort()
		return
	}

	c.Next()
}
