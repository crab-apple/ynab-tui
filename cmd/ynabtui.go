package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"log/slog"
	"os"
	"ynabtui/internal/app"
	"ynabtui/internal/files"
	"ynabtui/internal/ynabapi"
	"ynabtui/internal/ynabmodel"
	"ynabtui/test"
)

func main() {

	ynab := test.NewFakeYnab()

	ynab.SetTransactions([]ynabmodel.Transaction{
		test.MakeTransaction(&test.AccChecking, &test.CatGroceries, "2020-01-01", 12340, "Last minute groceries"),
		test.MakeTransaction(&test.AccCash, &test.CatGroceries, "2020-01-02", 3500, "Chewing gum"),
		test.MakeTransaction(&test.AccChecking, &test.CatRent, "2020-01-02", 1000000, ""),
	})

	runApp(os.Stdin, os.Stdout, ynab.Api(), files.AppFilesImpl{})
}

func runApp(input io.Reader, output io.Writer, api ynabapi.YnabApi, appFiles files.AppFiles) {

	defer setUpLogging(appFiles)()

	p := tea.NewProgram(app.InitialModel(api), tea.WithInput(input), tea.WithOutput(output))
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
