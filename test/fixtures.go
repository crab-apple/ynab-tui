package test

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"ynabtui/internal/ynabmodel"
	"ynabtui/internal/ynabmodel/date"
)

var (
	AccCash = ynabmodel.Account{
		Id:   uuid.UUID{},
		Name: "Cash",
	}
	AccChecking = ynabmodel.Account{
		Id:   uuid.UUID{},
		Name: "Checking account",
	}
)

var (
	CatGroceries = ynabmodel.Category{
		Id:   uuid.UUID{},
		Name: "Groceries",
	}
	CatRent = ynabmodel.Category{
		Id:   uuid.UUID{},
		Name: "Rent",
	}
)

func MakeTransaction(account *ynabmodel.Account, category *ynabmodel.Category, dateStr string, amount int64, memo string) ynabmodel.Transaction {

	d, err := date.Parse(dateStr)
	if err != nil {
		panic(err)
	}

	amountMoney, err := ynabmodel.NewMoney(amount)
	if err != nil {
		panic(err)
	}

	return ynabmodel.Transaction{
		Id:           fmt.Sprintf("%d", rand.Uint32()),
		Date:         d,
		AccountId:    account.Id,
		AccountName:  account.Name,
		CategoryId:   &category.Id,
		CategoryName: &category.Name,
		Amount:       amountMoney,
		Memo:         memo,
	}
}
