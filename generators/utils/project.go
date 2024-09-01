package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/rapidstellar/gohexa/configs"
)

// func CreateProject(name string, template string) {
// 	templateDir := "templates/" + template
// 	// templateDir := filepath.Join("templates", template)

// 	err := filepath.WalkDir(templateDir, func(path string, d fs.DirEntry, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		relPath, _ := filepath.Rel(templateDir, path)

// 		newPath := filepath.Join(name, relPath)

// 		if d.IsDir() {
// 			return os.MkdirAll(newPath, 0755)
// 		}

// 		// Copy file and replace placeholders
// 		content, err := os.ReadFile(path)
// 		if err != nil {
// 			return err
// 		}
// 		newContent := strings.ReplaceAll(string(content), "go-template", name)
// 		return os.WriteFile(newPath, []byte(newContent), 0644)
// 	})

// 	if err != nil {
// 		fmt.Printf("Error creating project: %v\n", err)
// 	} else {
// 		fmt.Printf("Project '%s' initialized successfully using the '%s' template!\n", name, template)
// 	}
// }

// FetchTemplateFromGitHub downloads and extracts a zip file from GitHub.
func FetchTemplateFromGitHub(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download template: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download template: HTTP status %s", resp.Status)
	}

	// Create a temporary file to store the downloaded zip file
	tmpFile, err := os.CreateTemp("", "template.zip")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy the response body to the temporary file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save downloaded zip file: %v", err)
	}

	// Open the zip file
	zipReader, err := zip.OpenReader(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	// Extract the zip file
	for _, file := range zipReader.File {
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in zip: %v", err)
		}
		defer rc.Close()

		// Determine the new file path
		newPath := filepath.Join(dest, file.Name)

		// Create directories if needed
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(newPath, 0755)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
			continue
		}

		// Create the file
		newFile, err := os.Create(newPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer newFile.Close()

		// Copy the file contents
		_, err = io.Copy(newFile, rc)
		if err != nil {
			return fmt.Errorf("failed to copy file contents: %v", err)
		}
	}

	return nil
}

// CreateProject initializes a new project using a template from GitHub.
func CreateProject(name, templateName string) {
	tmpDir := "tmp_template"
	templateURL := configs.TEMPLATE_URL
	// Fetch and extract the template
	err := FetchTemplateFromGitHub(templateURL, tmpDir)
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
