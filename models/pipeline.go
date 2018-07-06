package models

//go:generate reform

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/xaionaro-go/errors"
	"github.com/xaionaro-go/extime"
	"github.com/xaionaro-go/log"
	cfg "gitlab.telemed.help/devops/ci/config"
	"gitlab.telemed.help/devops/ci/gitlab"
	"gitlab.telemed.help/devops/ci/smtp"
	"golang.org/x/crypto/bcrypt"
)

//reform:pipelines
type Pipeline struct {
	Id               int          `reform:"id,pk"`
	GitlabPipelineId int          `reform:"gitlab_pipeline_id"`
	GitlabNamespace  string       `reform:"gitlab_namespace"`
	TokenHash        []byte       `reform:"token_hash"`
	ProjectName      string       `reform:"project_name" sql_size:"255"`
	TagName          string       `reform:"tag_name" sql_size:"255"`
	Success          *bool        `reform:"success"`
	CreatedAt        *extime.Time `reform:"created_at"`
	ApprovedAt       *extime.Time `reform:"approved_at"`
	DeletedAt        *extime.Time `reform:"deleted_at"`

	Approvals         Approvals         `reform:"-"`
	RequiredApprovals RequiredApprovals `reform:"-"`
}

type Pipelines []Pipeline

func (pipelines Pipelines) PrepareApprovals() Pipelines {
	ids := []int{}
	pipelineMap := map[int]*Pipeline{}
	for idx, pipeline := range pipelines {
		ids = append(ids, pipeline.Id)
		pipelineMap[pipeline.Id] = &pipelines[idx]
	}

	if len(ids) == 0 {
		return pipelines
	}

	approvals, err := ApprovalSQL.Select("`pipeline_id` IN (?)", ids)
	if err != nil {
		panic(err)
	}

	for _, approval := range approvals {
		pipelineMap[approval.PipelineId].Approvals = append(pipelineMap[approval.PipelineId].Approvals, approval)
	}

	return pipelines
}

func (pipelines Pipelines) PrepareRequiredApprovals() Pipelines {
	projectNames := []string{}
	pipelineMap := map[string][]*Pipeline{}
	for idx, pipeline := range pipelines {
		projectNames = append(projectNames, pipeline.ProjectName)
		pipelineMap[pipeline.ProjectName] = append(pipelineMap[pipeline.ProjectName], &pipelines[idx])
	}

	if len(projectNames) == 0 {
		return pipelines
	}

	requiredApprovals, err := RequiredApprovalSQL.Select("`project_name` IN (?)", projectNames)
	if err != nil {
		panic(err)
	}

	for _, requiredApproval := range requiredApprovals {
		for _, pipeline := range pipelineMap[requiredApproval.ProjectName] {
			pipeline.RequiredApprovals = append(pipeline.RequiredApprovals, requiredApproval)
		}
	}

	return pipelines
}

func (pipeline *Pipeline) SetToken(token string) *Pipeline {
	var err error
	pipeline.TokenHash, err = bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return pipeline
}

func (pipeline *Pipeline) CompareToken(token string) bool {
	return bcrypt.CompareHashAndPassword(pipeline.TokenHash, []byte(token)) == nil
}

func (pipeline *Pipeline) HideTokenHash() *Pipeline {
	pipeline.TokenHash = []byte("<hidden>")
	return pipeline
}

func (pipelines Pipelines) HideTokenHash() Pipelines {
	for idx := range pipelines {
		pipelines[idx].HideTokenHash()
	}
	return pipelines
}

func (pipeline Pipeline) GetUnsatisfiedGroups() (result []int) {
	requiredApprovals, err := RequiredApprovalSQL.Select(RequiredApproval{ProjectName: pipeline.ProjectName})
	if err == sql.ErrNoRows {
		return
	}
	if err != nil {
		panic(fmt.Errorf(`Cannot fetch current required approvals: %v`, err.Error()))
	}

	satisfiedApprovementGroupMap := map[int]bool{}
	usernameToApprovementGroupIdMap := map[string]int{}
	for _, requiredApproval := range requiredApprovals {
		satisfiedApprovementGroupMap[requiredApproval.ApprovementGroupId] = false
		usernameToApprovementGroupIdMap[strings.ToLower(requiredApproval.Username)] = requiredApproval.ApprovementGroupId
	}

	approvals, err := ApprovalSQL.Select(Approval{PipelineId: pipeline.Id})
	if err != nil && err != sql.ErrNoRows {
		panic(fmt.Errorf(`Cannot fetch current approvals: %v`, err.Error()))
	}

	for _, approval := range approvals {
		approvementGroupId := usernameToApprovementGroupIdMap[strings.ToLower(approval.Username)]
		satisfiedApprovementGroupMap[approvementGroupId] = true
	}

	for approvementGroupId, isApproved := range satisfiedApprovementGroupMap {
		if !isApproved {
			result = append(result, approvementGroupId)
		}
	}

	return
}

func (pipeline Pipeline) GetCommiterEmail() (string, error) {
	tag, err := gitlab.GetTag(pipeline.GitlabNamespace+"/"+pipeline.ProjectName, pipeline.TagName)
	if err != nil {
		return "", errors.CannotGetInfo.New(err, fmt.Sprintf("Cannot get tag description"), pipeline)
	}

	if tag.Commit == nil {
		return "", errors.OutOfRange.New(nil, "tag.Commit == nil", pipeline)
	}

	return tag.Commit.CommitterEmail, nil
}

func (pipeline Pipeline) AskAuthorForDescription() error {
	title := fmt.Sprintf("Please provide the release description for %v/%v",
		pipeline.ProjectName,
		pipeline.TagName,
	)
	message := title + ".\n\n" + fmt.Sprintf("URL: %v/%v/%v/tags/%v/release/edit", cfg.Get().GitLab.URL, pipeline.GitlabNamespace, pipeline.ProjectName, pipeline.TagName)

	email, err := pipeline.GetCommiterEmail()
	if err != nil {
		return err
	}

	return smtp.Send(email, title, message)
}

func (pipeline Pipeline) WaitForDescription() error {
	for {
		time.Sleep(time.Second * 5)

		tag, err := gitlab.GetTag(pipeline.GitlabNamespace+"/"+pipeline.ProjectName, pipeline.TagName)
		if err != nil {
			return errors.CannotGetInfo.New(err, fmt.Sprintf("Cannot get tag description"), pipeline)
		}

		if tag.Release != nil {
			if tag.Release.Description != "" {
				return nil
			}
		}
	}
}

func (pipeline Pipeline) AskApproversForApprovals() error {
	unsatisfiedGroups := pipeline.GetUnsatisfiedGroups()
	if len(unsatisfiedGroups) == 0 {
		return nil
	}

	requiredApprovals, err := RequiredApprovalSQL.Where(RequiredApproval{ProjectName: pipeline.ProjectName}).Select("approvement_group_id IN (?)", unsatisfiedGroups)
	if err != nil {
		return err
	}

	tag, err := gitlab.GetTag(pipeline.GitlabNamespace+"/"+pipeline.ProjectName, pipeline.TagName)
	if err != nil {
		return errors.CannotGetInfo.New(err, fmt.Sprintf("Cannot get tag description"), pipeline)
	}

	tagDescription := ""
	if tag.Release != nil {
		tagDescription = tag.Release.Description
		log.Errorf("Cannot get the release description")
	}

	for _, requiredApproval := range requiredApprovals {
		user := GetUserByUsername(requiredApproval.Username)
		user.CreateApproveToken(pipeline, tagDescription)
	}

	return nil
}

func (pipeline Pipeline) NotifyEverybody(title, message string) error {
	// Not implemented, yet
	return fmt.Errorf("Not implemented, yet")
}

func (pipeline *Pipeline) Approve() error {
	pipeline.ApprovedAt = &[]extime.Time{extime.Now()}[0]
	err := pipeline.Update()
	if err != nil {
		return fmt.Errorf(`Cannot update the pipeline info: %v`, err.Error())
	}

	message := fmt.Sprintf(`Deployment of %v/%v has been approved`, pipeline.ProjectName, pipeline.TagName)
	pipeline.NotifyEverybody(message, message)
	return nil
}
