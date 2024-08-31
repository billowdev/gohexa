package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	projectName := flag.String("name", "", "The name of the project")
	templateName := flag.String("template", "hexagonal", "The name of the template (default: hexagonal)")
	flag.Parse()

	if *projectName == "" {
		fmt.Println("Please provide a project name using the -name flag.")
		return
	}

	createProject(*projectName, *templateName)
}

func createProject(name string, template string) {
	templateDir := template
	err := filepath.WalkDir(templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(templateDir, path)
		newPath := filepath.Join(name, relPath)

		if d.IsDir() {
			return os.MkdirAll(newPath, 0755)
		}

		// Copy file and replace placeholders
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
		fmt.Printf("Project '%s' initialized successfully using the '%s' template!\n", name, template)
	}
}
