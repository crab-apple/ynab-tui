package ynabapi

import (
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type YnabApi interface {
	ReadTransactions(budgetId string, since date.Date) ([]ynabmodel.Transaction, error)
}
