package ynabapi

import (
	"github.com/google/uuid"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type YnabApi interface {
	ReadBudgets() ([]ynabmodel.Budget, error)
	ReadTransactions(budgetId uuid.UUID, since date.Date) ([]ynabmodel.Transaction, error)
}
