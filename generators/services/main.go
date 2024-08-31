// generators services
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
	featureName := flag.String("feature", "", "The name of the feature to generate a service for (e.g., Order)")
	outputDir := flag.String("output", "", "The output directory for the generated service files")
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

	generateServiceFile(*outputDir, *featureName, *projectName)
}


var ServiceTemplate = `
package services

import (
	"context"

	"github.com/{{ .ProjectName }}/internal/adapters/database"
	domain "github.com/{{ .ProjectName }}/internal/core/domain/{{ .FeatureName | ToLower }}"
	ports "github.com/{{ .ProjectName }}/internal/core/ports/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/pkg/configs"
	"github.com/{{ .ProjectName }}/pkg/helpers/pagination"
	"github.com/{{ .ProjectName }}/pkg/utils"
)

type {{ .FeatureName }}ServiceImpl struct {
	repo       ports.I{{ .FeatureName }}Repository
	transactor database.IDatabaseTransactor
}

func New{{ .FeatureName }}Service(
	repo ports.I{{ .FeatureName }}Repository,
	transactor database.IDatabaseTransactor,
) ports.I{{ .FeatureName }}Service {
	return &{{ .FeatureName }}ServiceImpl{repo: repo, transactor: transactor}
}

// Create{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Create{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}) utils.APIResponse {
	data := domain.To{{ .FeatureName }}Model(payload)
	if err := s.repo.Create{{ .FeatureName }}(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// Delete{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Delete{{ .FeatureName }}(ctx context.Context, id uint) utils.APIResponse {
	if err := s.repo.Delete{{ .FeatureName }}(ctx, id); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: nil}
}

// Get{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Get{{ .FeatureName }}(ctx context.Context, id uint) utils.APIResponse {
	data, err := s.repo.Get{{ .FeatureName }}(ctx, id)
	if err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	if data == nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Not Found", Data: nil}
	}
	res := domain.To{{ .FeatureName }}Domain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}

// Get{{ .FeatureName }}s implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Get{{ .FeatureName }}s(ctx context.Context) pagination.Pagination[[]domain.{{ .FeatureName }}] {
	data, err := s.repo.Get{{ .FeatureName }}s(ctx)
	if err != nil {
		return pagination.Pagination[[]domain.{{ .FeatureName }}]{}
	}
	// Convert repository data to domain models
	newData := utils.ConvertSlice(data.Rows, domain.To{{ .FeatureName }}Domain)
	return pagination.Pagination[[]domain.{{ .FeatureName }}]{
		Rows:       newData,
		Links:      data.Links,
		Total:      data.Total,
		Page:       data.Page,
		PageSize:   data.PageSize,
		TotalPages: data.TotalPages,
	}
}

// Update{{ .FeatureName }} implements ports.I{{ .FeatureName }}Service.
func (s *{{ .FeatureName }}ServiceImpl) Update{{ .FeatureName }}(ctx context.Context, payload domain.{{ .FeatureName }}) utils.APIResponse {
	data := domain.To{{ .FeatureName }}Model(payload)
	if err := s.repo.Update{{ .FeatureName }}(ctx, data); err != nil {
		return utils.APIResponse{StatusCode: configs.API_ERROR_CODE, StatusMessage: "Error", Data: err}
	}
	res := domain.To{{ .FeatureName }}Domain(data)
	return utils.APIResponse{StatusCode: configs.API_SUCCESS_CODE, StatusMessage: "Success", Data: res}
}
`

func generateServiceFile(dir string, featureName string, projectName string) {
	// Define the template for the service file
	

	// Prepare the data for template rendering
	data := struct {
		FeatureName string
		ProjectName string
	}{
		FeatureName: featureName,
		ProjectName: projectName,
	}

	// Parse and execute the template
	tmpl, err := template.New("service").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(ServiceTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_service.go", strings.ToLower(featureName))
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
		fmt.Printf("Service file '%s' created successfully!\n", filePath)
	}
}
