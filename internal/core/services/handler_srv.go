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

// ToLower returns the lowercase version of the input string
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Pluralize returns the plural form of the input string (simple example)
func Pluralize(s string) string {
	if strings.HasSuffix(s, "s") {
		return s
	}
	return s + "s"
}

// GenerateHandlerFile implements ports.IGeneratorService.
func (g *GeneratorServiceImpls) GenerateHandlerFile(dir string) {
	// Default to current directory if not provided
	defaultDir := "./internal/adapters/http/handlers"
	err := utils.EnsureDir(dir, defaultDir)
	if err != nil {
		fmt.Printf("Failed to ensure directory: %v", err)
	}

	// Prepare the data for template rendering
	data := domain.HandlerFlagDomain{
		FeatureName: g.flag.FeatureName,
		ProjectName: g.flag.ProjectName,
	}

	// Parse and execute the template
	tmpl, err := template.New("handler").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(domain.HandlerTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_handlers.go", strings.ToLower(g.flag.FeatureName))
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
		fmt.Printf("Handlers file '%s' created successfully!\n", filePath)
	}
}
