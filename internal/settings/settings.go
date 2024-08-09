package settings

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func ReadAccessToken() (string, error) {
	c, err := readYnabConfigFile("access_token")
	if err != nil {
		return "", fmt.Errorf("unable to read access token: %w", err)
	}
	return strings.TrimSpace(c), nil
}

type settings struct {
	BudgetId string `yaml:"budgetId"`
}

func ReadDefaultBudgetId() (string, error) {
	c, err := readYnabConfigFile("settings.yaml")
	if err != nil {
		return "", fmt.Errorf("unable to read settings: %w", err)
	}

	s := settings{}
	err = yaml.Unmarshal([]byte(c), &s)
	if err != nil {
		return "", fmt.Errorf("unable to read settings: %w", err)
	}

	return s.BudgetId, nil
}

func readYnabConfigFile(filename string) (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to read file: %w", err)
	}
	f := filepath.Join(homedir, ".ynab", filename)
	contents, err := os.ReadFile(f)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %w", err)
	}
	return string(contents), nil
}
