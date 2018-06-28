package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/markbates/goth"
	"github.com/sethvargo/go-password/password"
	"github.com/xaionaro-go/extime"
	cfg "gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/sms"
)

type User struct {
	GitLabUser goth.User
}

func GetUserByUsername(username string) *User {
	return &User{
		GitLabUser: goth.User{
			NickName: strings.ToLower(username),
		},
	}
}

func (user User) GetUsername() string {
	return strings.ToLower(user.GitLabUser.NickName)
}

func (user *User) CreateApproveToken(pipeline Pipeline) error {
	randomValue, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		return err
	}

	token := ApproveToken{Token: randomValue, Username: user.GetUsername(), PipelineId: pipeline.Id, CreatedAt: &[]extime.Time{extime.Now()}[0]}
	err = token.Create()
	if err != nil {
		return err
	}

	title := fmt.Sprintf("Please consider deployment of %v/%v",
		pipeline.ProjectName,
		pipeline.TagName,
	)
	message := title + fmt.Sprintf(" using URL: %v", cfg.Get().BaseURL+"/simpleApi/approveUsingToken/"+randomValue)
	user.SendNotification(title, message)
	return nil
}

func (user *User) ApprovePipeline(pipelineId int) error {
	err := (&Approval{
		Username:   user.GetUsername(),
		PipelineId: pipelineId,
		CreatedAt:  &[]extime.Time{extime.Now()}[0],
	}).Create()
	if err != nil {
		return fmt.Errorf(`Cannot create the approval: %v`, err.Error())
	}

	pipeline, err := PipelineSQL.First(Pipeline{Id: pipelineId})
	if err != nil {
		return fmt.Errorf(`Cannot fetch the pipeline info: %v`, err.Error())
	}

	unsatisfiedGroups := pipeline.GetUnsatisfiedGroups()
	if len(unsatisfiedGroups) > 0 {
		reqApproval, err := RequiredApprovalSQL.First(RequiredApproval{ProjectName: pipeline.ProjectName, Username: user.GetUsername()})
		if err != nil {
			return fmt.Errorf(`Cannot find the required approval: %v`, err.Error())
		}

		neighborsInGroup, err := RequiredApprovalSQL.Select(RequiredApproval{ProjectName: pipeline.ProjectName, ApprovementGroupId: reqApproval.ApprovementGroupId})
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf(`Cannot find the neighbors in the group: %v`, err.Error())
		}

		for _, neighbor := range neighborsInGroup {
			notifableUser := GetUserByUsername(neighbor.Username)
			message := fmt.Sprintf("User %v has already approved the deployment of %v/%v", pipeline.ProjectName, pipeline.TagName)
			notifableUser.SendNotification(message, message)
		}
		return nil
	}

	return pipeline.Approve()
}

func (user User) GetProfile() *UserProfile {
	profile, err := UserProfileSQL.First(UserProfile{Username: user.GetUsername()})
	if err == sql.ErrNoRows {
		return &UserProfile{Username: user.GetUsername()}
	}
	if err != nil {
		panic(err)
	}
	return &profile
}

func (user User) SendNotification(title, message string) error {
	profile := user.GetProfile()
	if profile.SMSTo == "" {
		return nil
	}
	return sms.Send("CI", profile.SMSTo, message)
}
