package main

import (
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

	transactions, err := ynabclient.ReadTransactions(token, budgetId)
	if err != nil {
		panic(err)
	}

	println(budgets)
	println(transactions)
}
