package models

//go:generate reform

import (
	"github.com/xaionaro-go/extime"
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
