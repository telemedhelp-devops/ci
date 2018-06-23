package helpers

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	m "gitlab.telemed.help/devops/ci/models"
)

func GetMe(c *gin.Context) m.User {
	session := sessions.Default(c)
	userI := session.Get("User")
	switch user := userI.(type) {
	case m.User:
		return user
	case string:
		var userT m.User
		err := json.Unmarshal([]byte(user), &userT)
		if err != nil {
			panic(err)
		}
		return userT
	}
	return m.User{}
}
