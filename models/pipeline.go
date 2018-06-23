package models

//go:generate reform

import (
	"github.com/xaionaro-go/extime"
)

//reform:pipelines
type Pipeline struct {
	Id          int          `reform:"id,pk"`
	ProjectName string       `reform:"project_name" sql_size:"255"`
	TagName     string       `reform:"tag_name" sql_size:"255"`
	CreatedAt   *extime.Time `reform:"created_at"`
	ApprovedAt  *extime.Time `reform:"approved_at"`

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
	projectNames:= []string{}
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

