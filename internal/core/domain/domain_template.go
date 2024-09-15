package domain

type DomainFlagDomain struct {
	FeatureName string
	ProjectName string
	UseUUID     bool
	DefaultUUID string
}

var DomainTemplate = `
package domain

import (
	"time"

	"github.com/{{ .ProjectName }}/internal/adapters/database/models"
)

type {{ .FeatureName }}Domain struct {
{{ if .UseUUID }}
	ID                 string    ` + "`json:\"id\"`" + `
{{ else }}
	ID                 uint      ` + "`json:\"id\"`" + `
{{ end }}
	CreatedAt          time.Time ` + "`json:\"created_at\"`" + `
	UpdatedAt          time.Time ` + "`json:\"updated_at\"`" + `
	Field1      string    ` + "`json:\"field_1\" validate: \"required,max=50\"`" + `
	Field2      string    ` + "`json:\"field_2\" validate: \"max=50\"`" + `
}

func To{{ .FeatureName }}Domain(data *models.{{ .FeatureName }}) {{ .FeatureName }}Domain {
	if data == nil {
		return {{ .FeatureName }}Domain{}
	}

	return {{ .FeatureName }}Domain{
		ID:                 data.ID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		Field1:             data.Field1,
		Field2:             data.Field2,
	}
}

func To{{ .FeatureName }}Model(data {{ .FeatureName }}Domain) (*models.{{ .FeatureName }}, error) {
	// validate data

	// return models
	return &models.{{ .FeatureName }}{
		ID:                 data.ID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		Field1:             data.Field1,
		Field2:             data.Field2,
	}, nil
}

type Update{{ .FeatureName }}Domain struct {
{{ if .UseUUID }}
	ID                 string    ` + "`json:\"id,omitempty\"`" + `
{{ else }}
	ID                 uint      ` + "`json:\"id,omitempty\"`" + `
{{ end }}
	CreatedAt          time.Time ` + "`json:\"created_at,omitempty\"`" + `
	UpdatedAt          time.Time ` + "`json:\"updated_at,omitempty\"`" + `
	Field1      string    ` + "`json:\"field_1,omitempty\" validate: \"omitempty,max=50\"`" + `
	Field2      string    ` + "`json:\"field_2,omitempty\" validate: \"omitempty,max=50\"`" + `
}

func UpdateDomainTo{{ .FeatureName }}Model(data {{ .FeatureName }}Domain) (*models.{{ .FeatureName }}, error) {
	// validate data

	// return models
	return &models.{{ .FeatureName }}{
		ID:                 data.ID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		Field1:             data.Field1,
		Field2:             data.Field2,
	}, nil
}
`
