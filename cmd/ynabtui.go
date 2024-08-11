package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rickb777/date/v2"
	"io"
	"log/slog"
	"os"
	"ynabtui/internal/files"
	"ynabtui/internal/model"
	"ynabtui/internal/settings"
	"ynabtui/internal/ynabclient"
)

func main() {
	runApp(os.Stdin, os.Stdout)
}

func runApp(input io.Reader, output io.Writer) {

	defer setUpLogging()()

	p := tea.NewProgram(model.InitialModel(), tea.WithInput(input), tea.WithOutput(output))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func setUpLogging() func() {

	logLevel := slog.LevelWarn
	_, d := os.LookupEnv("YNAB_TUI_DEBUG")
	if d {
		logLevel = slog.LevelDebug
	}

	filePath, err := files.GetAppFile("log")
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{Level: logLevel})))

	return func() {
		f.Close()
	}
}

func fetchDataFromYnabExample() {

	token, err := settings.ReadAccessToken()
	if err != nil {
		panic(err)
	}

	client, err := ynabclient.NewClient("https://api.ynab.com/v1", token)
	if err != nil {
		panic(err)
	}

	budgetId, err := settings.ReadDefaultBudgetId()
	if err != nil {
		panic(err)
	}

	budgets, err := client.ReadBudgets()
	if err != nil {
		panic(err)
	}
	slog.Info("Fetched budgets", "budgets", budgets)

	// Only get transactions form the past couple days
	sinceDate := date.TodayUTC().AddDate(0, 0, -2)
	if err != nil {
		panic(err)
	}

	transactions, err := client.ReadTransactions(budgetId, sinceDate)
	if err != nil {
		panic(err)
	}
	slog.Info("Fetched transactions", "transactions", transactions)
}
