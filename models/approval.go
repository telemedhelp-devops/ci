package models

//go:generate reform

import (
	"github.com/xaionaro-go/extime"
)

//reform:approvals
type Approval struct {
	Id         int          `reform:"id,pk"`
	Username   string       `reform:"username" sql_size:"255"`
	PipelineId int          `reform:"pipeline_id"`
	CreatedAt  *extime.Time `reform:"created_at"`
}

type Approvals []Approval
