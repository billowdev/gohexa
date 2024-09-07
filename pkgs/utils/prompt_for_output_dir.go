package utils

import (
	"fmt"
	"strings"
)

// promptForOutputDir prompts the user to provide an output directory.
// It returns a boolean indicating whether the directory was set and handled successfully.
func PromptForOutputDir(outputDir *string) bool {
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
			return CheckAndCreateDir(outputDir)
		default:
			fmt.Println("Invalid response. Exiting.")
			return false
		}
	}
}
