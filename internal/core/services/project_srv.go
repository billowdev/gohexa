package services

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/rapidstellar/gohexa/pkgs/configs"
	"github.com/rapidstellar/gohexa/pkgs/utils"
)

// CreateProject implements ports.IGeneratorService.
func (g *GeneratorServiceImpls) CreateProject(name string, templateName string) {
	tmpDir := "tmp_template"
	templateURL := configs.TEMPLATE_URL
	// Fetch and extract the template
	err := utils.FetchTemplateFromGitHub(templateURL, tmpDir)
	if err != nil {
		fmt.Printf("Error fetching template: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)

	// Define the path of the specific template directory within the extracted ZIP
	templateDir := filepath.Join(tmpDir, templateName)

	// Ensure the template directory exists
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		fmt.Printf("Template '%s' not found in ZIP file.\n", templateName)
		return
	}

	// Create the project directory
	err = os.MkdirAll(name, 0755)
	if err != nil {
		fmt.Printf("Error creating project directory: %v\n", err)
		return
	}

	// Process the extracted template files
	err = filepath.WalkDir(templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(templateDir, path)
		newPath := filepath.Join(name, relPath)

		if d.IsDir() {
			return os.MkdirAll(newPath, 0755)
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		newContent := strings.ReplaceAll(string(content), "go-template", name)
		return os.WriteFile(newPath, []byte(newContent), 0644)
	})

	if err != nil {
		fmt.Printf("Error creating project: %v\n", err)
	} else {
		fmt.Printf("Project '%s' initialized successfully using the '%s' template!\n", name, templateName)
	}
}
