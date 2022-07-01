package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetApplicationDir returns a string representation of the home path for use with configuration/data storage needs
func GetApplicationDir() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	return path, nil
}

// GetConfigPath returns a string representation of the configuration's path
func GetConfigPath() (string, error) {
	dir, err := GetApplicationDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(dir, "eol-cli.yaml")

	return configPath, nil
}
