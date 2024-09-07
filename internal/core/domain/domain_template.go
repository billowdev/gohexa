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
	ID                 string    ` + "`gorm:\"type:uuid;primaryKey;default:uuid_generate_v4()\" json:\"id\"`" + `
{{ else }}
	ID                 uint      ` + "`gorm:\"primaryKey;autoIncrement\" json:\"id\"`" + `
{{ end }}
	CreatedAt          time.Time ` + "`json:\"created_at\" gorm:\"autoCreateTime\"`" + `
	UpdatedAt          time.Time ` + "`json:\"updated_at\" gorm:\"autoUpdateTime\"`" + `
	Field1      string    ` + "`json:\"field_1\"`" + `
	Field2      string    ` + "`json:\"field_2\"`" + `
}

func To{{ .FeatureName }}Domain(data *models.{{ .FeatureName }}) {{ .FeatureName }}Domain {
	if data == nil {
		return {{ .FeatureName }}Domain{
			{{ if .UseUUID }}
			ID: "{{ .DefaultUUID }}",
			{{ else }}
			ID: 0,
			{{ end }}
		}
	}

	return {{ .FeatureName }}Domain{
		ID:                 data.ID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		Field1:             data.Field1,
		Field2:             defaultStringIfEmpty(data.Field2, "No Field2"),
	}
}

func To{{ .FeatureName }}Model(data {{ .FeatureName }}Domain) *models.{{ .FeatureName }} {
	return &models.{{ .FeatureName }}{
		ID:                 data.ID,
		CreatedAt:          data.CreatedAt,
		UpdatedAt:          data.UpdatedAt,
		Field1:             data.Field1,
		Field2:             defaultStringIfEmpty(data.Field2, "No Field2"),
	}
}

// defaultStringIfEmpty returns the default value if the input string is empty
func defaultStringIfEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
`
