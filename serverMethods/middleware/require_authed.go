package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	m "gitlab.telemed.help/devops/ci/models"
	"net/http"
)

func RequireAuthed(c *gin.Context) {
	session := sessions.Default(c)
	userJson, ok := session.Get("User").(string)
	if !ok {
		fmt.Println("user", session.Get("User"))
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Authorization required",
		})
		return
	}

	var user m.User
	err := json.Unmarshal(([]byte)(userJson), &user)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"Error": "Authorization required (cannot unserialize user data): " + err.Error(),
		})
		return
	}

	c.Set("User", user)

	c.Next()
	return
}
