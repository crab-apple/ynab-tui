package test

import (
	"ynabtui/internal/ynabapi"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

type FakeYnab struct {
	transactions []ynabmodel.Transaction
}

func NewFakeYnab() *FakeYnab {
	return &FakeYnab{}
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

func (api fakeYnabApi) ReadTransactions(budgetId string, since date.Date) ([]ynabmodel.Transaction, error) {
	return api.fy.transactions, nil
}
