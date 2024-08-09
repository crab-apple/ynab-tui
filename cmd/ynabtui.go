package main

import (
	"ynabtui/internal/ynabclient"
)

func main() {

	token, err := ynabclient.ReadAccessToken()
	if err != nil {
		panic(err)
	}

	budgets, err := ynabclient.ReadBudgets(token)
	if err != nil {
		panic(err)
	}

	println(budgets)
}
