package models

//go:generate reform

import (
	"github.com/xaionaro-go/extime"
)

//reform:approve_tokens
type ApproveToken struct {
	Id         int          `reform:"id,pk"`
	Token      string       `reform:"token" sql_size:"255"`
	Username   string       `reform:"username" sql_size:"255"`
	PipelineId int          `reform:"pipeline_id"`
	CreatedAt  *extime.Time `reform:"created_at"`
}
