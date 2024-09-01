package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rapidstellar/gohexa/generators/utils"
)

func main() {
	generateType := flag.String("generate", "", "Type of code to generate (options: project, transactor, model, domain, port, repository, service, handler, route, app)")
	projectName := flag.String("project", "my_project", "The name of the project (default: my_project)")
	featureName := flag.String("feature", "", "The name of the feature (required for app, domain, and model)")
	outputDir := flag.String("output", "", "The output directory for the generated files")
	templateName := flag.String("template", "hexagonal", "The name of the template (default: hexagonal)")
	useUUID := flag.Bool("uuid", false, "Use UUID for ID field instead of uint")
	help := flag.Bool("help", false, "Show help message")

	flag.Parse()

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
		utils.CreateProject(*outputDir, *templateName)
	case "transactor":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return
		}
		utils.GenerateTransactorFile(*outputDir)
	case "model":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		utils.GenerateModelsFile(*outputDir, *featureName, *projectName, *useUUID)

	case "domain":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		utils.GenerateDomainFile(*outputDir, *featureName, *projectName, *useUUID)
	case "port":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		utils.GeneratePortsFile(*outputDir, *featureName, *projectName)
	case "repository":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		utils.GenerateRepoFile(*outputDir, *featureName, *projectName)
	case "service":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		utils.GenerateServiceFile(*outputDir, *featureName, *projectName)
	case "handler":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		utils.GenerateHandlerFile(*outputDir, *featureName, *projectName)
	case "route":
		if *outputDir == "" {
			if !promptForOutputDir(outputDir) {
				return
			}
		}
		if *featureName == "" {
			fmt.Println("Please provide a feature name using -feature flags.")
			return
		}
		utils.GenerateRouteFile(*outputDir, *featureName, *projectName)
	case "app":
		if *featureName == "" || *outputDir == "" {
			fmt.Println("Please provide a feature name and output directory using -feature and -output flags.")
			return
		}
		utils.GenerateAppFile(*outputDir, *featureName, *projectName)
	default:
		fmt.Println("Invalid generate type. Options are: project, transactor, model, domain, port, repository, service, handler, route, app.")
	}
}

// promptForOutputDir prompts the user to provide an output directory.
// It returns a boolean indicating whether the directory was set and handled successfully.
func promptForOutputDir(outputDir *string) bool {
	for {
		fmt.Print("Please provide an output directory using the -output flag. (Default: Y, No: provide again): ")
		var response string
		fmt.Scanln(&response)

		switch strings.ToLower(strings.TrimSpace(response)) {
		case "y", "yes", "1":
			*outputDir = "" // Set your default directory here
			fmt.Println("Using default directory:", *outputDir)
			return true
		case "no", "n", "0":
			fmt.Print("Please provide the output directory again: ")
			fmt.Scanln(outputDir)
			if *outputDir == "" {
				fmt.Println("No directory provided. Exiting.")
				return false
			}
			return checkAndCreateDir(outputDir)
		default:
			fmt.Println("Invalid response. Exiting.")
			return false
		}
	}
}

// checkAndCreateDir checks if the directory exists and creates it if it doesn't.
// It returns a boolean indicating whether the directory handling was successful.
func checkAndCreateDir(outputDir *string) bool {
	if _, err := os.Stat(*outputDir); os.IsNotExist(err) {
		// Ask the user for confirmation to create the directory
		fmt.Printf("Directory '%s' does not exist. Do you want to create it? (y/n): ", *outputDir)
		var response string
		fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response)) // Normalize response

		if response != "y" && response != "yes" && response != "1" {
			fmt.Println("Directory creation aborted.")
			return false
		}

		// Create the directory
		if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			return false
		}

		fmt.Println("Directory created successfully.")
		return true
	} else if err != nil {
		// Handle other errors that might occur
		fmt.Printf("Error checking directory: %v\n", err)
		return false
	}

	return true
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
	fmt.Println("  go run github.com/rapidstellar/gohexa -generate project -output myproject")
	fmt.Println("    Generates a new project in the 'myproject' directory using the default 'hexagonal' template.")
	fmt.Println()
	fmt.Println("  go run github.com/rapidstellar/gohexa -generate project -output myproject -template hexagonal")
	fmt.Println("    Generates a new project in the 'myproject' directory using the 'hexagonal' template.")
	fmt.Println()
	fmt.Println("  go run github.com/rapidstellar/gohexa -generate model -output myproject -feature user -project my_project")
	fmt.Println("    Generates a model file for the 'user' feature in the 'myproject' directory.")
	fmt.Println()
	fmt.Println("  go run github.com/rapidstellar/gohexa -generate app -output myproject -feature user")
	fmt.Println("    Generates an app file for the 'user' feature in the 'myproject' directory.")
	fmt.Println()
	fmt.Println("For more information, visit the documentation at https://github.com/rapidstellar/gohexa")
}
