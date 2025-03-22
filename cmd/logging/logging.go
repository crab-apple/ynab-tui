package logging

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

func SetUpLogging() func() {

	logLevel := slog.LevelWarn
	_, d := os.LookupEnv("YNAB_TUI_DEBUG")
	if d {
		logLevel = slog.LevelDebug
	}

	logWriter, cleanup, err := getLogWriter()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: logLevel})))

	return cleanup
}

func getLogWriter() (io.Writer, func(), error) {
	filePath, err := getLogFilePath()
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

func getLogFilePath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homedir, ".ynab", "log"), nil
}
