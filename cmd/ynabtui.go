package main

import (
	"github.com/rickb777/date/v2"
	"ynabtui/internal/settings"
	"ynabtui/internal/ynabclient"
)

func main() {

	budgetId, err := settings.ReadDefaultBudgetId()
	if err != nil {
		panic(err)
	}

	token, err := settings.ReadAccessToken()
	if err != nil {
		panic(err)
	}

	budgets, err := ynabclient.ReadBudgets(token)
	if err != nil {
		panic(err)
	}

	// Only get transactions form the past couple days
	sinceDate := date.TodayUTC().AddDate(0, 0, -2)
	if err != nil {
		panic(err)
	}

	transactions, err := ynabclient.ReadTransactions(token, budgetId, sinceDate)
	if err != nil {
		panic(err)
	}

	println(budgets)
	println(transactions)
}
