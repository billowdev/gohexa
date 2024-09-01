package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func GeneratePortsFile(dir string, featureName string, projectName string) {
	// Default to current directory if not provided
	if dir == "" {
		dir = "./internal/core/ports"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return
		}
	}
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
