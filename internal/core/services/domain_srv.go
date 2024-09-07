package services

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/rapidstellar/gohexa/internal/core/domain"
	"github.com/rapidstellar/gohexa/pkgs/utils"
)

// GenerateDomainFile implements ports.IGeneratorService.
func (g *GeneratorServiceImpls) GenerateDomainFile(dir string, useUUID bool) {
	// Define the template for the domain file
	defaultDir := "./internal/core/domain"
	err := utils.EnsureDir(dir, defaultDir)
	if err != nil {
		fmt.Printf("Failed to ensure directory: %v", err)
	}
	// Prepare the data for template rendering
	data := domain.DomainFlagDomain{
		FeatureName: g.flag.FeatureName,
		ProjectName: g.flag.ProjectName,
		UseUUID:     useUUID,
		DefaultUUID: "00000000-0000-0000-0000-000000000000", // Default UUID value
	}

	// Parse and execute the template
	tmpl, err := template.New("domain").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(domain.DomainTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_domain.go", strings.ToLower(g.flag.FeatureName))
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
