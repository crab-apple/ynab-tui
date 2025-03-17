package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAppFile(filename string) (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".ynab", filename), nil
}

func ReadYnabConfigFile(filename string) (string, error) {
	f, err := GetAppFile(filename)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %w", err)
	}
	contents, err := os.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %w", err)
	}
	return string(contents), nil
}
