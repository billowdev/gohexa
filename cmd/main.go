package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rapidstellar/gohexa/generators/utils"
)

func main() {
	generateType := flag.String("generate", "", "Type of code to generate (options: project, model, domain, app, transactor, repository, service, handler, route)")
	projectName := flag.String("project", "my_project", "The name of the project (default: my_project)")
	featureName := flag.String("feature", "", "The name of the feature (required for app, domain, and model)")
	outputDir := flag.String("output", "", "The output directory for the generated files")
	templateName := flag.String("template", "hexagonal", "The name of the template (default: hexagonal)")
	useUUID := flag.Bool("uuid", false, "Use UUID for ID field instead of uint")
	flag.Parse()

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
		utils.GenerateTransactorFile(*outputDir)
	case "app":
		if *featureName == "" || *outputDir == "" {
			fmt.Println("Please provide a feature name and output directory using -feature and -output flags.")
			return
		}
		utils.GenerateAppFile(*outputDir, *featureName, *projectName)
	case "domain":
		if *featureName == "" || *outputDir == "" {
			fmt.Println("Please provide a feature name and output directory using -feature and -output flags.")
			return
		}
		utils.GenerateDomainFile(*outputDir, *featureName, *projectName, *useUUID)
	case "model":
		if *featureName == "" || *outputDir == "" {
			fmt.Println("Please provide a feature name and output directory using -feature and -output flags.")
			return
		}
		utils.GenerateModelsFile(*outputDir, *featureName, *projectName, *useUUID)
	case "port":

	case "repository":
		if *featureName == "" {
			fmt.Println("Please provide a feature name using the -feature flag.")
			return
		}
		utils.GenerateRepoFile(*outputDir, *featureName, *projectName)
	case "service":
		if *featureName == "" {
			fmt.Println("Please provide a feature name using the -feature flag.")
			return
		}
		if *outputDir == "" {
			fmt.Println("Please provide an output directory using the -output flag.")
			return
		}
		utils.GenerateServiceFile(*outputDir, *featureName, *projectName)
	case "handler":
		if *featureName == "" {
			fmt.Println("Please provide a feature name using the -feature flag.")
			return
		}

		if *outputDir == "" {
			fmt.Println("Please provide an output directory using the -output flag.")
			return
		}
		utils.GenerateHandlerFile(*outputDir, *featureName, *projectName)
	case "route":
		if *featureName == "" {
			fmt.Println("Please provide a feature name using the -feature flag.")
			return
		}
		if *outputDir == "" {
			fmt.Println("Please provide an output directory using the -output flag.")
			return
		}
		utils.GenerateRouteFile(*outputDir, *featureName, *projectName)
	default:
		fmt.Println("Invalid generate type. Options are: project, app, domain, model.")
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
