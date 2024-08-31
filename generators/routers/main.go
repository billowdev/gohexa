package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	
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

func main() {
	featureName := flag.String("feature", "", "The name of the feature to generate routes for (e.g., SeaPort)")
	outputDir := flag.String("output", "", "The output directory for the generated route files")
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

	generateRouteFile(*outputDir, *featureName, *projectName)
}

var RouteTemplate = `
package routers

import (
	handlers "{{ .ProjectName }}/internal/adapters/handlers/{{ .FeatureName | ToLower }}"
	"{{ .ProjectName }}/pkg/middlewares"
)

func (r RouterImpl) Create{{ .FeatureName }}Routes(h handlers.I{{ .FeatureName }}Handler) {
	r.route.Get("/{{ .FeatureName | Pluralize | ToLower }}", h.HandleGet{{ .FeatureName }}s)
	r.route.Get("/{{ .FeatureName | Pluralize | ToLower }}/:id", h.HandleGet{{ .FeatureName }})
	r.route.Post("/{{ .FeatureName | Pluralize | ToLower }}", h.HandleCreate{{ .FeatureName }})
	r.route.Put("/{{ .FeatureName | Pluralize | ToLower }}/:id", h.HandleUpdate{{ .FeatureName }})
	r.route.Delete("/{{ .FeatureName | Pluralize | ToLower }}/:id", h.HandleDelete{{ .FeatureName }})
}
`

func generateRouteFile(dir string, featureName string, projectName string) {
	// Prepare the data for template rendering
	data := struct {
		FeatureName string
		ProjectName string
	}{
		FeatureName: featureName,
		ProjectName: projectName,
	}

	// Parse and execute the template
	tmpl, err := template.New("route").Funcs(template.FuncMap{
		"ToLower":  strings.ToLower,
		"Pluralize": func(s string) string {
			if strings.HasSuffix(s, "s") {
				return s
			}
			return s + "s"
		},
	}).Parse(RouteTemplate)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Create the output file path
	fileName := fmt.Sprintf("%s_routes.go", strings.ToLower(featureName))
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
		fmt.Printf("Route file '%s' created successfully!\n", filePath)
	}
}
