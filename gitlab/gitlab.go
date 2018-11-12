package gitlab

import (
	"fmt"
	"strings"

	"github.com/xaionaro-go/errors"
	"github.com/xanzy/go-gitlab"
	cfg "gitlab.telemed.help/devops/ci/config"
)

type (
	Tag           = gitlab.Tag
	User          = gitlab.User
	Job           = gitlab.Job
	ErrorResponse = gitlab.ErrorResponse
)

var (
	client *gitlab.Client
)

func init() {
	cfg.AddReloadHook(reloadConfig)
}

func reloadConfig() {
	c := cfg.Get()
	client = gitlab.NewClient(nil, c.GitLab.Token)
	client.SetBaseURL(c.GitLab.URL + "/api/v4")
}

func GetTag(projectId interface{}, tagName string) (*Tag, error) {
	tag, _, err := client.Tags.GetTag(projectId, tagName)
	return tag, err
}

func GetUserByName(username string) (*User, error) {
	users, _, err := client.Users.ListUsers(&gitlab.ListUsersOptions{Username: &username})
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if strings.ToLower(user.Username) == username {
			return user, nil
		}
	}

	return nil, errors.NotFound.New(nil, fmt.Sprintf("User with username %v is not found", username), User{}, username)
}

func GetJob(projectId int, jobId int) (*Job, error) {
	job, _, err := client.Jobs.GetJob(projectId, jobId)
	return job, err
}
