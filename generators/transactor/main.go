// init_transactor.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rapidstellar/gohexa/internal/core/domain"
	"github.com/rapidstellar/gohexa/internal/core/services"
)

func main() {
	outputDir := flag.String("output", "", "The output directory for the generated transactor file")
	flag.Parse()

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
		FeatureName: "",
		ProjectName: "",
	})
	srv.GenerateTransactorFile(*outputDir)
}
