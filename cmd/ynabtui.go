package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/rickb777/date/v2"
	"log/slog"
	"os"
	"ynabtui/internal/model"
	"ynabtui/internal/settings"
	"ynabtui/internal/ynabclient"
)

func main() {
	p := tea.NewProgram(model.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
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
