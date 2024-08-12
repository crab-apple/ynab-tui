package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"log/slog"
	"os"
	"ynabtui/internal/files"
	"ynabtui/internal/model"
)

func main() {
	runApp(os.Stdin, os.Stdout, files.AppFilesImpl{})
}

func runApp(input io.Reader, output io.Writer, appFiles files.AppFiles) {

	defer setUpLogging(appFiles)()

	p := tea.NewProgram(model.InitialModel(), tea.WithInput(input), tea.WithOutput(output))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func setUpLogging(appFiles files.AppFiles) func() {

	logLevel := slog.LevelWarn
	_, d := os.LookupEnv("YNAB_TUI_DEBUG")
	if d {
		logLevel = slog.LevelDebug
	}

	logWriter, cleanup, err := appFiles.GetLogWriter()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: logLevel})))

	return cleanup
}
