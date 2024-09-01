package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)



func GenerateDomainFile(dir string, featureName string, projectName string, useUUID bool) {
	// Define the template for the domain file
	if dir == "" {
		dir = "./internal/core/domain"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return
		}
	}
	// Prepare the data for template rendering
	data := struct {
		FeatureName string
		ProjectName string
		UseUUID     bool
		DefaultUUID string
	}{
		FeatureName: featureName,
		ProjectName: projectName,
		UseUUID:     useUUID,
		DefaultUUID: "00000000-0000-0000-0000-000000000000", // Default UUID value
	}

	// Parse and execute the template
	tmpl, err := template.New("domain").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(DomainTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_domain.go", strings.ToLower(featureName))
	filePath := filepath.Join(dir, fileName)

	// Create the output file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Execute the template and write to the file
	err = tmpl.Execute(file, data)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	} else {
		fmt.Printf("Domain file '%s' created successfully!\n", filePath)
	}
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