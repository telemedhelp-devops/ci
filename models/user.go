package models

import (
	"github.com/markbates/goth"
)

type User struct {
	GitLabUser goth.User
}
