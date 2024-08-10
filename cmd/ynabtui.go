package main

import (
	"github.com/rickb777/date/v2"
	"ynabtui/internal/settings"
	"ynabtui/internal/ynabclient"
)

func main() {

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

	// Only get transactions form the past couple days
	sinceDate := date.TodayUTC().AddDate(0, 0, -2)
	if err != nil {
		panic(err)
	}

	transactions, err := client.ReadTransactions(budgetId, sinceDate)
	if err != nil {
		panic(err)
	}

	println(budgets)
	println(transactions)
}
