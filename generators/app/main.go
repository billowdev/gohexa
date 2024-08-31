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
	featureName := flag.String("feature", "", "The name of the feature to generate an app file for (e.g., SystemField)")
	outputDir := flag.String("output", "", "The output directory for the generated app file")
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

	generateAppFile(*outputDir, *featureName, *projectName)
}

var AppTemplate = `
package app

import (
	"github.com/{{ .ProjectName }}/internal/adapters/database"
	handlers "github.com/{{ .ProjectName }}/internal/adapters/http/handlers/{{ .FeatureName | ToLower }}"
	"github.com/{{ .ProjectName }}/internal/adapters/http/routers"
	repositories "github.com/{{ .ProjectName }}/internal/adapters/repositories/{{ .FeatureName | ToLower }}"
	services "github.com/{{ .ProjectName }}/internal/core/services/{{ .FeatureName | ToLower }}"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppContainer(app *fiber.App, db *gorm.DB) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	{{ .FeatureName }}App(route, db)
	return app
}

func {{ .FeatureName }}App(r routers.RouterImpl, db *gorm.DB) {
	transactorRepo := database.NewTransactorRepo(db)
	{{ .FeatureName | ToLower }}Repo := repositories.New{{ .FeatureName }}Repo(db)
	{{ .FeatureName | ToLower }}Srv := services.New{{ .FeatureName }}Service({{ .FeatureName | ToLower }}Repo, transactorRepo)
	{{ .FeatureName | ToLower }}Handlers := handlers.New{{ .FeatureName }}Handler({{ .FeatureName | ToLower }}Srv)
	r.Create{{ .FeatureName }}Route({{ .FeatureName | ToLower }}Handlers)
}
`

func generateAppFile(dir string, featureName string, projectName string) {
	// Define the template for the app file

	// Prepare the data for template rendering
	data := struct {
		FeatureName string
		ProjectName string
	}{
		FeatureName: featureName,
		ProjectName: projectName,
	}

	// Parse and execute the template
	tmpl, err := template.New("app").Funcs(template.FuncMap{
		"ToLower": strings.ToLower,
	}).Parse(AppTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_app.go", strings.ToLower(featureName))
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
		fmt.Printf("App file '%s' created successfully!\n", filePath)
	}
}
