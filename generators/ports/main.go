// generators ports
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
	featureName := flag.String("feature", "", "The name of the feature to generate ports for (e.g., Order)")
	outputDir := flag.String("output", "", "The output directory for the generated port files")
	projectName := flag.String("project", "my_project", "The name of the project (default: my_project)")
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

	generatePortsFile(*outputDir, *featureName, *projectName)
}

var PortsTemplate = `
package ports

import (
	"context"

	"github.com/{{ .ProjectName }}/internal/adapters/database/models"
	domain "github.com/{{ .ProjectName }}/internal/core/domain/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/pkg/helpers/pagination"
	"github.com/{{ .ProjectName }}/pkg/utils"
)

type I{{ .FeatureName }}Repository interface {
	Get{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) (*models.{{ .FeatureName }}, error)
	Get{{ .FeatureName }}s(ctx context.Context) (*pagination.Pagination[[]models.{{ .FeatureName }}], error)
	Create{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error
	Update{{ .FeatureName }}(ctx context.Context, payload *models.{{ .FeatureName }}) error
	Delete{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) error
}

type I{{ .FeatureName }}Service interface {
	Get{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) utils.APIResponse
	Get{{ .FeatureName }}s(ctx context.Context) pagination.Pagination[[]domain.{{ .FeatureName }}Domain]
	Create{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}Domain) utils.APIResponse
	Update{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}Domain) utils.APIResponse
	Delete{{ .FeatureName }}(ctx context.Context, id {{ .IDType }}) utils.APIResponse
}
`
func generatePortsFile(dir string, featureName string, projectName string) {
	// Define the template for the ports file
	

	// Prepare the data for template rendering
	data := struct {
		FeatureName string
		ProjectName string
		IDType      string
	}{
		FeatureName: featureName,
		ProjectName: projectName,
		IDType:      "uint", // Default to uint, can be changed to string if UUID is used
	}

	// Parse and execute the template
	tmpl, err := template.New("ports").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(PortsTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_ports.go", strings.ToLower(featureName))
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
		fmt.Printf("Ports file '%s' created successfully!\n", filePath)
	}
}
