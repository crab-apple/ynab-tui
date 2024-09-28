package ynabapi

import (
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type YnabApi interface {
	ReadBudgets() ([]ynabmodel.Budget, error)
	ReadTransactions(budgetId string, since date.Date) ([]ynabmodel.Transaction, error)
}
