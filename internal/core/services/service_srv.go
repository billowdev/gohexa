package services

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/rapidstellar/gohexa/internal/core/domain"
)

// GenerateServiceFile implements ports.IGeneratorService.
func (g *GeneratorServiceImpls) GenerateServiceFile(dir string) {
	// Define the template for the service file
	if dir == "" {
		dir = "./internal/core/services"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return
		}
	}

	// Prepare the data for template rendering
	data := domain.ServiceFlagDomain{
		FeatureName: g.flag.FeatureName,
		ProjectName: g.flag.ProjectName,
	}

	// Parse and execute the template
	tmpl, err := template.New("service").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(domain.ServiceTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_service.go", strings.ToLower(g.flag.FeatureName))
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
