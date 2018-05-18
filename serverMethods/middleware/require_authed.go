package middleware

import (
	"github.com/gin-gonic/gin"
)

func init() {
}

func RequireAuthed(c *gin.Context) {
	c.Next()
}
