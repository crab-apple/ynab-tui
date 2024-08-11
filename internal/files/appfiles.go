package files

import (
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
