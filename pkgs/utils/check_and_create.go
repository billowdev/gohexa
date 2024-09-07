package utils

import (
	"fmt"
	"os"
	"strings"
)

// checkAndCreateDir checks if the directory exists and creates it if it doesn't.
// It returns a boolean indicating whether the directory handling was successful.
func CheckAndCreateDir(outputDir *string) bool {
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
