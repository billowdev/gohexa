package adapters

import (
	"fmt"
	"os"

	"github.com/rapidstellar/gohexa/internal/core/domain"
	"github.com/rapidstellar/gohexa/internal/core/services"
	"github.com/rapidstellar/gohexa/pkgs/utils"
)

type IGeneratorAdapter interface {
	GohexaGeneratorAdapter(flag domain.GeneratorFlag)
}

type GenratorAdapter struct{}

func NewGeneratorAdapter() IGeneratorAdapter {
	return &GenratorAdapter{}
}

// GohexaGeneratorAdapter implements IGeneratorAdapter.
func (g *GenratorAdapter) GohexaGeneratorAdapter(gf domain.GeneratorFlag) {
	featureName := gf.FeatureName
	projectName := gf.ProjectName
	generateType := gf.GenerateType
	outputDir := gf.OutputDir
	templateName := gf.TemplateName
	useUUID := gf.UseUUID
	help := gf.Help

	srv := services.NewGeneratorService(domain.GeneratorFlagDomain{
		FeatureName: *featureName,
		ProjectName: *projectName,
	})

	if *help {
		showHelp()
		return
	}

	if *generateType == "" {
		fmt.Println("Please specify a generate type using the -generate flag.")
		return
	}

	// // Ensure the output directory exists, create if not
	// if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
	// 	fmt.Printf("Error creating directories: %v\n", err)
	// 	return
	// }
	switch *generateType {
	case "project":
		if *outputDir == "" {
			fmt.Println("Please provide an output directory using the -output flag.")
			return
		}
		srv.CreateProject(*outputDir, *templateName)
	case "transactor":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return
		}
		srv.GenerateTransactorFile(*outputDir)
	case "model":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		srv.GenerateModelsFile(*outputDir, *useUUID)

	case "domain":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		srv.GenerateDomainFile(*outputDir, *useUUID)
	case "port":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		srv.GeneratePortsFile(*outputDir)
	case "repository":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		srv.GenerateRepoFile(*outputDir)
	case "service":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		srv.GenerateServiceFile(*outputDir)
	case "handler":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		srv.GenerateHandlerFile(*outputDir)
	case "route":
		if *outputDir == "" {
			if !utils.PromptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		srv.GenerateRouteFile(*outputDir)
	case "app":
		if *featureName == "" || *outputDir == "" {
			fmt.Println("Please provide a feature name and output directory using -feature and -output flags.")
			return
		}
		srv.GenerateAppFile(*outputDir)
	default:
		fmt.Println("Invalid generate type. Options are: project, transactor, model, domain, port, repository, service, handler, route, app.")
	}

}

// showHelp displays the help message for the command-line tool
func showHelp() {
	fmt.Println("Usage: gohexa [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -generate string   Type of code to generate. Options include:")
	fmt.Println("                      project        - Generates a new project structure.")
	fmt.Println("                      transactor     - Generates a transactor file.")
	fmt.Println("                      model          - Generates a model file. Requires -feature flag.")
	fmt.Println("                      domain         - Generates a domain file. Requires -feature flag.")
	fmt.Println("                      port           - Generates a port file. Requires -feature flag.")
	fmt.Println("                      repository     - Generates a repository file. Requires -feature flag.")
	fmt.Println("                      service        - Generates a service file. Requires -feature flag.")
	fmt.Println("                      handler        - Generates a handler file. Requires -feature flag.")
	fmt.Println("                      route          - Generates a route file. Requires -feature flag.")
	fmt.Println("                      app            - Generates an app file. Requires -feature and -output flags.")
	fmt.Println()
	fmt.Println("  -project string    The name of the project to generate. Default is 'my_project'.")
	fmt.Println("  -feature string    The name of the feature for which to generate files. Required for:")
	fmt.Println("                      model, domain, port, repository, service, handler, route, app")
	fmt.Println()
	fmt.Println("  -output string     The directory where the generated files will be placed. This flag is required for:")
	fmt.Println("                      project, transactor, model, domain, port, repository, service, handler, route, app")
	fmt.Println("                    If not provided, you will be prompted to enter it.")
	fmt.Println()
	fmt.Println("  -template string   The template to use for generating the project. Default is 'hexagonal'.")
	fmt.Println("                    Examples: 'hexagonal', 'hexa-fiber', 'hexa-grpc'")
	fmt.Println()
	fmt.Println("  -uuid              Use UUID for ID fields instead of uint. Default is false.")
	fmt.Println()
	fmt.Println("  -help              Show this help message and exit.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  gohexa -generate project -output myproject")
	fmt.Println("    Generates a new project in the 'myproject' directory using the default 'hexagonal' template.")
	fmt.Println()
	fmt.Println("  gohexa -generate project -output myproject -template hexagonal")
	fmt.Println("    Generates a new project in the 'myproject' directory using the 'hexagonal' template.")
	fmt.Println()
	fmt.Println("  gohexa -generate model -output myproject -feature user -project my_project")
	fmt.Println("    Generates a model file for the 'user' feature in the 'myproject' directory.")
	fmt.Println()
	fmt.Println("  gohexa -generate app -output myproject -feature user")
	fmt.Println("    Generates an app file for the 'user' feature in the 'myproject' directory.")
	fmt.Println()
	fmt.Println("For more information, visit the documentation at https://github.com/rapidstellar/gohexa")
}
