package test

import (
	"github.com/google/uuid"
	"time"
	ynabmodel2 "ynabtui/app/app/ynabmodel"
	"ynabtui/app/app/ynabmodel/date"
	"ynabtui/app/driven_ports"
)

type FakeYnab struct {
	budgets      []ynabmodel2.Budget
	transactions []ynabmodel2.Transaction
}

func NewFakeYnab() *FakeYnab {
	return &FakeYnab{}
}

func (fy *FakeYnab) SetBudgets(budgets []ynabmodel2.Budget) {
	fy.budgets = budgets
}
func (fy *FakeYnab) SetTransactions(transactions []ynabmodel2.Transaction) {
	fy.transactions = transactions
}

func (fy *FakeYnab) Api() driven_ports.ForCommunicatingWithYnab {
	return fakeYnabApi{fy: fy}
}

type fakeYnabApi struct {
	fy *FakeYnab
}

func (api fakeYnabApi) ReadBudgets() ([]ynabmodel2.Budget, error) {
	return api.fy.budgets, nil
}

func (api fakeYnabApi) ReadTransactions(budgetId uuid.UUID, since date.Date) ([]ynabmodel2.Transaction, error) {
	time.Sleep(10 * time.Millisecond)
	return api.fy.transactions, nil
}
