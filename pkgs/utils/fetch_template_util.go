package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// FetchTemplateFromGitHub downloads and extracts a zip file from GitHub.
func FetchTemplateFromGitHub(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download template: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download template: HTTP status %s", resp.Status)
	}

	// Create a temporary file to store the downloaded zip file
	tmpFile, err := os.CreateTemp("", "template.zip")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy the response body to the temporary file
	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save downloaded zip file: %v", err)
	}

	// Open the zip file
	zipReader, err := zip.OpenReader(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	// Extract the zip file
	for _, file := range zipReader.File {
		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in zip: %v", err)
		}
		defer rc.Close()

		// Determine the new file path
		newPath := filepath.Join(dest, file.Name)

		// Create directories if needed
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(newPath, 0755)
			if err != nil {
				return fmt.Errorf("failed to create directory: %v", err)
			}
			continue
		}

		// Create the file
		newFile, err := os.Create(newPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer newFile.Close()

		// Copy the file contents
		_, err = io.Copy(newFile, rc)
		if err != nil {
			return fmt.Errorf("failed to copy file contents: %v", err)
		}
	}

	return nil
}
