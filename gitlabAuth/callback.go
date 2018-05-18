package gitlabAuth

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	m "gitlab.telemed.help/devops/ci/models"
	"net/http"
)

func Callback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		fmt.Fprintln(c.Writer, err)
		return
	}

	session := sessions.Default(c)
	session.Set("User", m.User{GitLabUser: user})
	session.Save()

	c.Redirect(http.StatusFound, "/")
	return
}
