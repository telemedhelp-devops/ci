package gitlabAuth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Login(c *gin.Context) {
	ctx := context.WithValue(c.Request.Context(), "provider", "gitlab")
	gothic.BeginAuthHandler(c.Writer, c.Request.WithContext(ctx))
}
