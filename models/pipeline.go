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
}
