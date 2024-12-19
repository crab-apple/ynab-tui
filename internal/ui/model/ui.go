package model

import (
	"github.com/samber/lo"
	"ynabtui/internal/ynabapi"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type UI struct {
	api ynabapi.YnabApi
}

func NewUI(api ynabapi.YnabApi) UI {
	return UI{api: api}
}

func (ui UI) FirstLoad() Screen {
	since, _ := date.Today().MinusDays(7)

	budgets, err := ui.api.ReadBudgets()
	if err != nil {
		// TODO handle
		panic(err)
	}

	budget := lo.MaxBy(budgets, func(a ynabmodel.Budget, b ynabmodel.Budget) bool {
		return a.LastModifiedOn.After(b.LastModifiedOn)
	})

	transactions, err := ui.api.ReadTransactions(budget.Id, since)
	if err != nil {
		// TODO handle
		panic(err)
	}

	screen := TransactionsScreen{
		Transactions: transactions,
	}
	return screen
}
