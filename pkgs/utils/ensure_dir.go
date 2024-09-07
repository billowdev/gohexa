package utils

import (
	"fmt"
	"os"
)

// EnsureDir ensures that the given directory exists, creating it if necessary.
// If dir is empty, it uses the defaultDir instead.
func EnsureDir(dir, defaultDir string) error {
	if dir == "" {
		dir = defaultDir
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating directories: %v", err)
	}
	return nil
}
