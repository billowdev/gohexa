package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func GenerateModelsFile(dir string, featureName string, projectName string, useUUID bool) {
	// Default to current directory if not provided

	if dir == "" {
		dir = "./internal/adapters/database/models"
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
	}{
		FeatureName: featureName,
		ProjectName: projectName,
		UseUUID:     useUUID,
	}

	// Parse and execute the template
	tmpl, err := template.New("models").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(ModelsTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s.go", strings.ToLower(featureName))
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
		fmt.Printf("Model file '%s' created successfully!\n", filePath)
	}
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
