package files

import (
	"io"
	"os"
	"path/filepath"
)

func GetLogWriter() (io.Writer, func(), error) {
	filePath, err := GetAppFile("log")
	if err != nil {
		return nil, nil, err
	}
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	cleanup := func() { f.Close() }

	return f, cleanup, nil
}

func GetAppFile(filename string) (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".ynab", filename), nil
}
