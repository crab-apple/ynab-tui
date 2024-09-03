package test

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
	"ynabtui/internal/model/ynab"
)

var (
	AccCash = ynab.Account{
		Id:   uuid.UUID{},
		Name: "Cash",
	}
	AccChecking = ynab.Account{
		Id:   uuid.UUID{},
		Name: "Checking account",
	}
)

var (
	CatGroceries = ynab.Category{
		Id:   uuid.UUID{},
		Name: "Groceries",
	}
	CatRent = ynab.Category{
		Id:   uuid.UUID{},
		Name: "Rent",
	}
)

func MakeTransaction(account *ynab.Account, category *ynab.Category, date string, amount int64, memo string) ynab.Transaction {

	t, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", date))
	if err != nil {
		panic(err)
	}

	amountMoney, err := ynab.NewMoney(amount)
	if err != nil {
		panic(err)
	}

	return ynab.Transaction{
		Id:           fmt.Sprintf("%d", rand.Uint32()),
		Date:         t,
		AccountId:    account.Id,
		AccountName:  account.Name,
		CategoryId:   &category.Id,
		CategoryName: &category.Name,
		Amount:       amountMoney,
		Memo:         memo,
	}
}
