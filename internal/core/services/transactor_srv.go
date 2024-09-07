package services

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/rapidstellar/gohexa/internal/core/domain"
	"github.com/rapidstellar/gohexa/pkgs/utils"
)

// GenerateTransactorFile implements ports.IGeneratorService.
func (g *GeneratorServiceImpls) GenerateTransactorFile(dir string) {
	// Define the template for the transactor file
	defaultDir := "./internal/adapters/database"
	err := utils.EnsureDir(dir, defaultDir)
	if err != nil {
		fmt.Printf("Failed to ensure directory: %v", err)
	}
	// Create the output file path
	fileName := "transactor.go"
	filePath := filepath.Join(dir, fileName)

	// Create the output file
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Parse and execute the template
	tmpl, err := template.New("transactor").Parse(domain.TransactorTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	err = tmpl.Execute(file, nil)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	} else {
		fmt.Printf("Transactor file '%s' created successfully!\n", filePath)
	}
}
