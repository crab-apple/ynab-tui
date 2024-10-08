package test

import (
	"github.com/google/uuid"
	"time"
	"ynabtui/internal/ynabapi"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type FakeYnab struct {
	budgets      []ynabmodel.Budget
	transactions []ynabmodel.Transaction
}

func NewFakeYnab() *FakeYnab {
	return &FakeYnab{}
}

func (fy *FakeYnab) SetBudgets(budgets []ynabmodel.Budget) {
	fy.budgets = budgets
}
func (fy *FakeYnab) SetTransactions(transactions []ynabmodel.Transaction) {
	fy.transactions = transactions
}

func (fy *FakeYnab) Api() ynabapi.YnabApi {
	return fakeYnabApi{fy: fy}
}

type fakeYnabApi struct {
	fy *FakeYnab
}

func (api fakeYnabApi) ReadBudgets() ([]ynabmodel.Budget, error) {
	return api.fy.budgets, nil
}

func (api fakeYnabApi) ReadTransactions(budgetId uuid.UUID, since date.Date) ([]ynabmodel.Transaction, error) {
	time.Sleep(10 * time.Millisecond)
	return api.fy.transactions, nil
}
