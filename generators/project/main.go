package main

import (
	"flag"
	"fmt"

	"github.com/rapidstellar/gohexa/internal/core/domain"
	"github.com/rapidstellar/gohexa/internal/core/services"
)

func main() {
	projectName := flag.String("name", "", "The name of the project")
	templateName := flag.String("template", "hexagonal", "The name of the template (default: hexagonal)")
	flag.Parse()

	if *projectName == "" {
		fmt.Println("Please provide a project name using the -name flag.")
		return
	}
	srv := services.NewGeneratorService(domain.GeneratorFlagDomain{
		FeatureName: "",
		ProjectName: *projectName,
	})
	srv.CreateProject(*projectName, *templateName)
}
