package domain

type ModelFlagDomain struct {
	FeatureName string
	ProjectName string
	UseUUID     bool
}

var ModelsTemplate = `
package models

import (
	"time"

	"gorm.io/gorm"
)

type {{ .FeatureName }} struct {
	gorm.Model
	{{ if .UseUUID }}ID                 string         ` + "`gorm:\"type:uuid;primaryKey;default:uuid_generate_v4()\" json:\"id\"`" + `{{ else }}ID                 uint           ` + "`gorm:\"primaryKey;autoIncrement\" json:\"id\"`" + `{{ end }}
	CreatedAt          time.Time      ` + "`json:\"created_at\" gorm:\"autoCreateTime\"`" + `
	UpdatedAt          time.Time      ` + "`json:\"updated_at\" gorm:\"autoUpdateTime\"`" + `
	DeletedAt          gorm.DeletedAt ` + "`gorm:\"index\" json:\"deleted_at,omitempty\"`" + `
}

var TN{{ .FeatureName }} = "{{ .FeatureName | ToLower }}s"

func (st *{{ .FeatureName }}) TableName() string {
	return TN{{ .FeatureName }}
}
`
