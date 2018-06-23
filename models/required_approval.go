package models

//go:generate reform

import (
	"github.com/xaionaro-go/extime"
)

//reform:required_approvals
type RequiredApproval struct {
	Id          int          `reform:"id,pk"`
	Username    string       `reform:"username" sql_size:"255"`
	ProjectName string       `reform:"project_name" sql_size:"255"`
	CreatedAt   *extime.Time `reform:"created_at"`
}

type RequiredApprovals []RequiredApproval
