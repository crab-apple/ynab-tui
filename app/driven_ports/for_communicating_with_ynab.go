package driven_ports

import (
	"github.com/google/uuid"
	"ynabtui/app/app/ynabmodel"
	"ynabtui/app/app/ynabmodel/date"
)

type ForCommunicatingWithYnab interface {
	ReadBudgets() ([]ynabmodel.Budget, error)
	ReadTransactions(budgetId uuid.UUID, since date.Date) ([]ynabmodel.Transaction, error)
}
