package main

import (
	"flag"

	adapters "github.com/rapidstellar/gohexa/internal/adapters/generators"
	"github.com/rapidstellar/gohexa/internal/core/domain"
)

func main() {
	generateType := flag.String("generate", "", "Type of code to generate (options: project, transactor, model, domain, port, repository, service, handler, route, app)")
	projectName := flag.String("project", "my_project", "The name of the project (default: my_project)")
	featureName := flag.String("feature", "", "The name of the feature Example Order, Document")
	outputDir := flag.String("output", "", "The output directory for the generated files")
	templateName := flag.String("template", "hexagonal", "The name of the template (default: hexagonal)")
	useUUID := flag.Bool("uuid", false, "Use UUID for ID field instead of uint")
	help := flag.Bool("help", false, "Show help message")
	flag.Parse()

	generatorFlag := domain.GeneratorFlag{
		GenerateType: generateType,
		ProjectName:  projectName,
		FeatureName:  featureName,
		OutputDir:    outputDir,
		TemplateName: templateName,
		UseUUID:      useUUID,
		Help:         help,
	}
	genrator := adapters.NewGeneratorAdapter()
	genrator.GohexaGeneratorAdapter(generatorFlag)

}
