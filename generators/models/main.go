// generators models
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rapidstellar/gohexa/internal/core/domain"
	"github.com/rapidstellar/gohexa/internal/core/services"
)

func main() {
	featureName := flag.String("feature", "", "The name of the feature to generate models for (e.g., Order)")
	outputDir := flag.String("output", "", "The output directory for the generated model files")
	projectName := flag.String("project", "my_project", "The name of the project (default: my_project)")
	useUUID := flag.Bool("uuid", false, "Whether to use UUID for the ID field (default: false)")
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
	srv := services.NewGeneratorService(domain.GeneratorFlagDomain{
		FeatureName: *featureName,
		ProjectName: *projectName,
	})
	srv.GenerateModelsFile(*outputDir, *useUUID)
}
