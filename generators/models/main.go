// generators models
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	featureName := flag.String("feature", "", "The name of the feature to generate models for (e.g., Order)")
	outputDir := flag.String("output", "", "The output directory for the generated model files")
	projectName := flag.String("project", "my_project", "The name of the project (default: my_project)")
	useUUID := flag.Bool("uuid", false, "Whether to use UUID for the ID field (default: false)")
	flag.Parse()

	if *featureName == "" {
		fmt.Println("Please provide a feature name using the -feature flag.")
		return
	}

	if *outputDir == "" {
		fmt.Println("Please provide an output directory using the -output flag.")
		return
	}

	// Ensure the output directory exists, create if not
	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	generateModelsFile(*outputDir, *featureName, *projectName, *useUUID)
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

func generateModelsFile(dir string, featureName string, projectName string, useUUID bool) {
	// Define the template for the model file


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
