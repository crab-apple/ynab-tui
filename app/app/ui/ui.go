package ui

import (
	"github.com/samber/lo"
	"ynabtui/app/app/ynabmodel"
	"ynabtui/app/app/ynabmodel/date"
	"ynabtui/app/driven_ports"
)

type UI struct {
	api driven_ports.ForCommunicatingWithYnab
}

func NewUI(api driven_ports.ForCommunicatingWithYnab) UI {
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
