package gitlabAuth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	m "gitlab.telemed.help/devops/ci/models"
	"net/http"
	"net/url"
)

func validateState(req *http.Request, sess goth.Session) error {
	rawAuthURL, err := sess.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	originalState := authURL.Query().Get("state")
	if originalState != "" && (originalState != req.URL.Query().Get("state")) {
		return errors.New("state token mismatch")
	}
	return nil
}

func Callback(c *gin.Context) {
	/*
		req := c.Request
		res := c.Writer

		defer gothic.Logout(res, req)

		provider, err := goth.GetProvider("gitlab")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_description": "invalid request (case #1)",
			})
		}

		value, err := gothic.GetFromSession("gitlab", req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_description": "invalid request (case #2)",
			})
			return
		}

		sess, err := provider.UnmarshalSession(value)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_description": "invalid request (case #3)",
			})
			return
		}

		err = validateState(req, sess)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_description": "invalid request (case #4)",
			})
			return
		}

		_, err = provider.FetchUser(sess)
		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_description": "already authed",
			})
			return
		}

		_, err = sess.Authorize(provider, req.URL.Query())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_description": "invalid request (case #5)",
			})
			return
		}

		gitlabUser, err := provider.FetchUser(sess)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error_description": "invalid request (case #6)",
			})
			return
		}
	*/
	ctx := context.WithValue(c.Request.Context(), "provider", "gitlab")
	gitlabUser, err := gothic.CompleteUserAuth(c.Writer, c.Request.WithContext(ctx))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ErrorDescription": "invalid request: " + err.Error(),
		})
		return
	}

	userBytes, err := json.Marshal(m.User{GitLabUser: gitlabUser})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"ErrorDescription": "cannot serialize user data",
		})
		return
	}
	userString := string(userBytes)
	session := sessions.Default(c)
	session.Set("User", userString)
	session.Save()

	c.Redirect(http.StatusFound, "/")
	return
}
