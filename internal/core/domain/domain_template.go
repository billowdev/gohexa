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
	Field1      string    ` + "`json:\"field_1\"`" + `
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
	}
}

func To{{ .FeatureName }}Model(data {{ .FeatureName }}Domain) *models.{{ .FeatureName }} {
	return &models.{{ .FeatureName }}{
		ID:                 data.ID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		Field1:             data.Field1,
	}
}
`
