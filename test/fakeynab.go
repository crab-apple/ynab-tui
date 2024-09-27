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
	return &FakeYnab{
		transactions: []ynabmodel.Transaction{
			MakeTransaction(&AccChecking, &CatGroceries, "2020-01-01", 12340, "Last minute groceries"),
			MakeTransaction(&AccCash, &CatGroceries, "2020-01-02", 3500, "Chewing gum"),
			MakeTransaction(&AccChecking, &CatRent, "2020-01-02", 1000000, ""),
		},
	}
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
